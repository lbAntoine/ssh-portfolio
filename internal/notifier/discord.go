package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type discordField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

type discordFooter struct {
	Text string `json:"text"`
}

type discordEmbed struct {
	Title     string         `json:"title"`
	Color     int            `json:"color"`
	Fields    []discordField `json:"fields"`
	Footer    discordFooter  `json:"footer"`
	Timestamp string         `json:"timestamp"`
}

type discordPayload struct {
	Embeds []discordEmbed `json:"embeds"`
}

// SendDiscord posts a contact form submission as a Discord embed.
func SendDiscord(webhookURL, name, email, message string) error {
	payload := discordPayload{
		Embeds: []discordEmbed{
			{
				Title: fmt.Sprintf("#%s from ssh-portfolio", name),
				Color: 0x6C91C2,
				Fields: []discordField{
					{Name: "name", Value: name, Inline: true},
					{Name: "email", Value: email, Inline: true},
					{Name: "message", Value: message},
				},
				Footer:    discordFooter{Text: "ssh antoinelb.fr -p 2222"},
				Timestamp: time.Now().UTC().Format(time.RFC3339),
			},
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	resp, err := http.Post(webhookURL, "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("discord webhook returned %d", resp.StatusCode)
	}
	return nil
}
