package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Config struct {
	OrgId     string
	ProjectId string
	ApiKey    string
}

func CreateGPtInstance(c Config) *OpenAi {
	return &OpenAi{c, "gpt-3.5-turbo", "https://api.openai.com/v1/chat/completions", 0.7}
}

type OpenAi struct {
	config      Config
	model       string
	url         string
	temperature float64
}

func (g *OpenAi) SetModel(model string) *OpenAi {
	g.model = model
	return g
}

func (g *OpenAi) SetTemperature(temperature float64) *OpenAi {
	g.temperature = temperature
	return g
}

func (g *OpenAi) Query(messages []Message) (*Message, error) {

	body := g.CreateQueryBody(messages)

	client := &http.Client{}

	req, _ := http.NewRequest("POST", g.url, bytes.NewReader(body))
	authheader := fmt.Sprintf("Bearer %v", g.config.ApiKey)

	req.Header.Set("Authorization", authheader)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	b, _ := io.ReadAll(resp.Body)

	var response GptResponse

	if err := json.Unmarshal(b, &response); err != nil {
		fmt.Println(err)

		return nil, err
	}
	return &response.Choices[0].Message, nil

}
