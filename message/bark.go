package message

import (
	"bytes"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"sendmsg/model"
)

type Bark struct {
	url string
	key string
}

func (b *Bark) Send(body Body) {
	log.Printf("[bark] sending message...")
	var reqBody = model.BarkRequest{
		DeviceKey: b.key,
		Title:     body.Title,
		Body:      body.Content,
	}
	req, _ := json.Marshal(reqBody)
	resp, err := http.Post(b.url,
		"application/json; charset=utf-8",
		bytes.NewReader(req))
	if err != nil {
		log.Fatalf("[bark] http post failed: %v\n", err)
	}
	defer resp.Body.Close()

	log.Printf("[bark] Send successful")
}
