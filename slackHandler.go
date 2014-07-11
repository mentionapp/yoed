package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type slackHandler struct {
	webhookUrl string
}

type slackPayload struct {
	Username  string `json:"username"`
	Text      string `json:"text"`
	IconEmoji string `json:"icon_emoji"`
}

func (h *slackHandler) Handle(username string) {

	payload := &slackPayload{
		Username:  username,
		Text:      ":yo:",
		IconEmoji: ":yo:",
	}

	buf, err := json.Marshal(payload)

	if err != nil {
		log.Printf("slackHandler: %s", err)
		return
	}

	resp, err := http.Post(h.webhookUrl, "application/json", bytes.NewBuffer(buf))

	if err != nil {
		log.Printf("slackHandler: %s", err)
		return
	}

	defer resp.Body.Close()

	log.Printf("slackHandler: %s", resp.Status)

	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		log.Printf("slackHandler: %s", err)
	} else {
		log.Printf("slackHandler: %s", string(body))
	}

}
