package message

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sendmsg/common"
	"sendmsg/model"
)

type Bark struct {
	url string
	key string
}

func (b *Bark) Send(body Body) {
	common.Log.Printf("[bark] sending message...")
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
		common.Log.Fatalf("[bark] http post failed: %v\n", err)
	}
	defer resp.Body.Close()

	common.Log.Printf("[bark] Send successful")
}
