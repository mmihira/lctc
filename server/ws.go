package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/jmoiron/sqlx"
	"io"
	models "lctc/models"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{} // use default options

type WsHandler struct {
	db *sqlx.DB
}

type Msg struct {
	Ids []string `json:"ids"`
}

func (h WsHandler) getTweetIds() []string {
	return models.TweetsSinceTimestamp(int32(time.Now().Unix())-20, h.db)
}

func (h WsHandler) index(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	done := make(chan struct{}, 1)
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	} // use default options
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	ticker := time.NewTicker(5 * time.Second)

	defer func() {
		log.Print("Closing")
		ticker.Stop()
		done <- struct{}{}
		conn.Close()
		log.Print("Closed")
	}()

	for {
		select {
		case <-ctx.Done():
			log.Print("Done")
			return
		case <-done:
			log.Print("Done")
			return
		case <-ticker.C:
			m := Msg{h.getTweetIds()}
			b, _ := json.Marshal(&m)

			w, werr := conn.NextWriter(websocket.TextMessage)
			if werr != nil {
				log.Print("error", err)
				return
			}
			_, err := io.WriteString(w, string(b))
			if err != nil {
				log.Print("error", err)
				return
			}
		}
	}

}
