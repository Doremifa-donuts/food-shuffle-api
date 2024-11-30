package ws

import "fmt"

// クライアント管理の構造体
type Hub struct {
	clients    map[string]*Client
	register   chan *Client
	unregister chan *Client
	share      chan []string
	boost      chan []string
}

// ハブを初期化する
func NewHub() *Hub {
	return &Hub{
		clients:    make(map[string]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		share:      make(chan []string),
		boost:      make(chan []string),
	}
}

// ハブの状態を管理する関数
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client.userUuid] = client
		case client := <-h.unregister:
			if _, ok := h.clients[client.userUuid]; ok {
				delete(h.clients, client.userUuid)
				close(client.send)
			}
		case shareUuids := <-h.share:
			fmt.Println("shareUuids", shareUuids)
			var keys []string
			for key := range h.clients {
				keys = append(keys, key)
			}
			fmt.Println("clients:", keys)
			for _, shareUuid := range shareUuids {
				h.clients[shareUuid].send <- []byte("新たなレビューを取得しました")
			}

		case boostUuids := <-h.boost:
			for _, boostUuid := range boostUuids {
				h.clients[boostUuid].send <- []byte("お助けブーストが届きました")
			}
		}
	}
}
