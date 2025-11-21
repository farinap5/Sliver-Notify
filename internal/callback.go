package internal

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func (c Config)SendData(d string) {
	c.setMessage(d)

	if c.Telegram.Token != "" && c.Telegram.Chat != "" {
		log.Println("Sending Telegram")
		c.SendTelegram()
	}

	if c.Discord.Token != "" && c.Discord.Chat != "" {
		log.Println("Sending Discord")
		c.SendDiscord()
	}

	if c.Teams.Webhook != "" {
		log.Println("Sending Teams")
		c.SendTeams()
	}
}

func (c *Config)setMessage(d string) {
	c.Message = d
}

func (c Config)SendTelegram() {
	url := "https://api.telegram.org/bot" + c.Telegram.Token + "/sendMessage"


	payload := map[string]string{
    	"text": c.Message,
		"chat_id": c.Telegram.Chat,
	}
	jsonStr, _ := json.Marshal(payload)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Println(err.Error())
    }
    defer resp.Body.Close()
    //body, _ := io.ReadAll(resp.Body)
}

func (c Config)SendDiscord() {
	url := "https://discord.com/api/webhooks/" + c.Discord.Chat + "/" + c.Discord.Token

	payload := map[string]string{
    	"content": c.Message,
	}
	jsonStr, _ := json.Marshal(payload)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Println(err.Error())
    }
    defer resp.Body.Close()
    //body, _ := io.ReadAll(resp.Body)
}

func (c Config)SendTeams() {
	url := c.Teams.Webhook

	payload := map[string]string{
    	"Rota": c.Message,
		"IP": "placeholder",
		"User-Agent":"placeholder",
	}
	jsonStr, _ := json.Marshal(payload)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Println(err.Error())
    }
    defer resp.Body.Close()
    //body, _ := io.ReadAll(resp.Body)
}