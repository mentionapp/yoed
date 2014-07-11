package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type yobackHandler struct {
	apiToken string
}

func (h *yobackHandler) Handle(username string) {

	resp, err := http.PostForm("http://api.justyo.co/yo/", url.Values{
		"api_token": {h.apiToken},
		"username":  {username},
	})

	if err != nil {
		log.Printf("yobackHandler: %s", err)
		return
	}

	defer resp.Body.Close()

	log.Printf("yobackHandler: %s", resp.Status)

	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		log.Printf("yobackHandler: %s", err)
	} else {
		log.Printf("yobackHandler: %s", string(body))
	}
}
