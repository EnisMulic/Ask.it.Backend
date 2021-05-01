package websockets

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/EnisMulic/Ask.it.Backend/contracts/responses"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type NotificationHandler struct {
}

func NewNotificationHandler() *NotificationHandler {
	return &NotificationHandler{}
}

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool { return true },
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return nil, err
    }

    return conn, nil
}

func (nc *NotificationHandler) ServeWS(pool *Pool, rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: "Unable to convert id",
		})

		out, _ := json.Marshal(errors)

		http.Error(rw, string(out), http.StatusNotFound)
		return
	}

	conn, err := Upgrade(rw, r)
    if err != nil {
        fmt.Fprintf(rw, "%+V\n", err)
		return
    }
    
	client := &Client{
		ID: uint64(id),
        Conn: conn,
        Pool: pool,
    }

    pool.Register <- client
    client.Read()
}
