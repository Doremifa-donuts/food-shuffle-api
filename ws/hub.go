package ws

import "fmt"

// クライアント管理の構造体
type Hub struct {
	clients    map[string]*Client
	register   chan *Client
	unregister chan *Client
	share      chan []string
	boost      chan map[string][]string
}

// ハブを初期化する
func NewHub() *Hub {
	return &Hub{
		clients:    make(map[string]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		share:      make(chan []string),
		boost:      make(chan map[string][]string),
	}
}

// ハブの状態を管理する関数
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register: // クライアント登録
			h.clients[client.userUuid] = client
		case client := <-h.unregister: // クライアント登録解除
			if _, ok := h.clients[client.userUuid]; ok {
				delete(h.clients, client.userUuid)
				close(client.send)
			}
		case shareUuids := <-h.share: // レビューの通知送信
			for _, shareUuid := range shareUuids {
				if client, ok := h.clients[shareUuid]; ok {
					client.send <- []byte("新たなレビューを取得しました")
				}
			}
		case boostContents := <-h.boost: // ブーストの送信
			for item, uuids := range boostContents {
				for _, uuid := range uuids {
					if client, ok := h.clients[uuid]; ok {
						// クライアントが存在する場合のみ送信
						client.send <- []byte(item + "からお助け要請が届きました")
					} else {
						fmt.Printf("クライアント %s が見つかりません\n", uuid)
					}
				}
			}
		}
	}
}

// ハブのブースト機能に内容を追加する
func SetBoost(boost map[string][]string) {
	hub.boost <- boost
}
