package ws

import (
	logging "food-shuffle-api/log"
	"food-shuffle-api/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// ヘルスチェック関連の設定
var (
	pongWait   = 30 * time.Second  // 周期を30秒に設定
	pingPeriod = pongWait * 9 / 10 // 周期より少し短い周期でpingを設定する
)

// WSの通信のサイズやタイムアウトの設定
var (
	writeWait      = 5 * time.Second
	maxMessageSize = int64(512) //メッセージサイズを最大512に制限
)

type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	userUuid string

	// Buffered channel of outbound messages.
	send chan []byte
}

var locationShareService service.LocationShareService

// クライアントにデータを送信する
func (c *Client) write() {

	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// クライアントからデータを取得する
func (c *Client) read() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	// メッセージを受信したらタイムアウトまでの経過時間をリセットする
	c.conn.SetPongHandler(
		func(string) error {
			c.conn.SetReadDeadline(time.Now().Add(pongWait))
			return nil
		})

	for {
		// wsから位置情報の受信をする
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				// 切断された時の処理
				logging.LogError("websocket read message is unexpected close error", err)
			}
			break
		}
		//
		userUuids, err := locationShareService.NotifyReviewByLocationMessage(c.userUuid, message)
		if err != nil {
			logging.LogError("review share service falied", nil)
		} else {
			SetReview(userUuids)
		}
	}
}

// 位置情報共有のWS通信を管理する
func LocationShareHandler(ctx *gin.Context) {
	var upgrader = websocket.Upgrader{
		WriteBufferSize: 1024,
		ReadBufferSize:  1024,
	}

	uuid, _ := ctx.Get("uuid")
	uuidString := uuid.(string)

	// WSにアップグレードする
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		logging.LogError("can not upgrade websocket", err)
		return
	}

	// クライアント構造体を作る
	client := &Client{
		hub:      hub,
		conn:     conn,
		userUuid: uuidString,
		send:     make(chan []byte, 256),
	}
	client.hub.register <- client

	// データ取得の処理を行う
	go client.write()
	go client.read()
}
