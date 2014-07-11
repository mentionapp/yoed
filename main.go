package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type yoedConfig struct {
	Listen   string `json:"listen"`
	Handlers struct {
		Slack *struct {
			WebhookUrl string `json:"webhook_url"`
		} `json:"slack"`
		Yoback *struct {
			ApiToken string `json:"api_token"`
		} `json:"yoback"`
	} `json:"handlers"`
}

type yoedHandler interface {
	Handle(username string)
}

func loadConfig(configPath string) (*yoedConfig, error) {

	configFile, err := os.Open(configPath)

	if err != nil {
		return nil, err
	}

	configJson, err := ioutil.ReadAll(configFile)

	if err != nil {
		return nil, err
	}

	config := &yoedConfig{}

	if err := json.Unmarshal(configJson, config); err != nil {
		return nil, err
	}

	return config, nil
}

func main() {

	config, err := loadConfig("./config.json")

	if err != nil {
		panic(fmt.Sprintf("failed loading config: %s", err))
	}

	handlers := make(map[string]yoedHandler)

	log.Println("Adding handlers")

	if slack := config.Handlers.Slack; slack != nil {
		log.Println("Adding slack handler")
		handlers["slack"] = &slackHandler{
			webhookUrl: slack.WebhookUrl,
		}
	}

	if yoback := config.Handlers.Yoback; yoback != nil {
		log.Println("Adding yoback handler")
		handlers["yoback"] = &yobackHandler{
			apiToken: yoback.ApiToken,
		}
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/yoed", func(w http.ResponseWriter, r *http.Request) {

		username := r.URL.Query().Get("username")
		log.Printf("got a YO from %s", username)

		for name, handler := range handlers {
			log.Printf("handling using %s for %s", name, username)
			handler.Handle(username)
		}
	})

	server := http.Server{
		Addr:    config.Listen,
		Handler: mux,
	}

	log.Printf("Listening...")

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}

}
