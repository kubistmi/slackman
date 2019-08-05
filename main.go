package slackman

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// The Message structure
type Message struct {
	API     string `json:"-"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
	User    string `json:"username"`
	Icon    string `json:"icon_url"`
	AsUser  bool   `json:"as_user"`
}

// NewMessage initialises a message
func NewMessage(API, channel, user, text, icon string) *Message {
	new := Message{
		API:     API,
		Channel: channel,
		Text:    text,
		User:    user,
		Icon:    icon,
		AsUser:  false}
	return (&new)
}

// Send sends the prepared message
func (m *Message) Send() string {

	load, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", "https://slack.com/api/chat.postMessage", bytes.NewBuffer(load))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+m.API)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		log.Fatalf("The http response code was %v\n", resp.StatusCode)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return string(bodyBytes)
}
