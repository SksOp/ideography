package openai

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type QueryBody struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
}

func (g *OpenAi) CreateUserMessageObject(q string) Message {
	return Message{Role: "user", Content: q}
}

func (g *OpenAi) CreateSystemMessageObject(q string) Message {
	return Message{Role: "system", Content: q}

}

func (g *OpenAi) CreateQueryBody(q []Message) []byte {
	qb := &QueryBody{
		Model:       g.model,
		Temperature: g.temperature,
		Messages:    q,
	}
	fmt.Println(qb)
	body, _ := json.Marshal(qb)
	return body
}
