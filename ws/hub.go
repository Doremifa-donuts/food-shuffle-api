package ws

import (
	"encoding/json"
	"fmt"
	logging "food-shuffle-api/log"
)

// クライアント管理の構造体
type Hub struct {
	clients    map[string]*Client
	register   chan *Client
	unregister chan *Client
	// share      chan []string
	// boost      chan map[string][]string
	// boost chan dto.BoostNotify
	notify chan NotifyProvider
}

// ハブを初期化する
func NewHub() *Hub {
	return &Hub{
		clients:    make(map[string]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		// share:      make(chan []string),
		// boost:      make(chan map[string][]string),
		// boost: make(chan dto.BoostNotify),
		notify: make(chan NotifyProvider),
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
		case notifyProvider := <-h.notify:
			// 通知内容を受け取る
			content := notifyProvider.Content
			// 通知タイプによって通知内容を作成する
			switch notifyProvider.Type {
			case Review:
				notifyProvider.Message = "新たなレビューを取得しました"
			case Boost:
				if content, ok := notifyProvider.Content.(BoostContent); ok {
					notifyProvider.Message = content.RestaurantName + "からお助け要請が届きました"
				} else {
					logging.LogError("Incorrect struct type passed as argument", nil)
				}
			}

			// 通知内容をjsonにエンコードする
			json, err := json.Marshal(content)
			if err != nil {
				logging.LogError("failed to marshal json for boost contents", err)
				return
			}
			// クライアントに対して通知を行う
			for _, uuid := range notifyProvider.UserUuids {
				if client, ok := h.clients[uuid]; ok {
					// クライアントが存在する場合のみ送信
					client.send <- []byte(json)
				} else {
					fmt.Printf("クライアント %s が見つかりません\n", uuid)
				}
			}
		}
	}
}

// ハブのブースト機能に内容を追加する
func SetBoost(userUuids []string, boostContent BoostContent) {
	hub.notify <- NotifyProvider{UserUuids: userUuids, Type: Boost, Content: boostContent}
}
func SetReview(userUuids []string) {
	hub.notify <- NotifyProvider{UserUuids: userUuids, Type: Review}
}
