package websockets

import "fmt"

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    []*Client
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
        Unregister: make(chan *Client),
		Clients:    []*Client{},
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
        select {
        case client := <-pool.Register:
            pool.Clients = append(pool.Clients, client)
        case client := <-pool.Unregister:
			for i, c := range pool.Clients {
				if client.ID == c.ID {
					pool.Clients = append(pool.Clients[:i], pool.Clients[i+1:]...)
					break
				}
			}
        case message := <-pool.Broadcast:
            fmt.Println("Sending message to client in Pool")
			for _, client := range pool.Clients {
				if client.ID == message.ClientID {
					if err := client.Conn.WriteJSON(message); err != nil {
						fmt.Println(err)
						return
					}
				}
				
			}
        }
    }
}