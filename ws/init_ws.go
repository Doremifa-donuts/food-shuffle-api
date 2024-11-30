package ws

var hub *Hub

func InitWebsocket() {
	hub = NewHub()
	go hub.Run()
}
