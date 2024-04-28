package llm

import (
	"ideography/packages/openai"
	"os"
)

type Message = openai.Message

type LLM interface {
	Query(messages []Message) (*Message, error)
	CreateUserMessageObject(q string) Message
	CreateSystemMessageObject(q string) Message
}

func NewLLM() LLM {
	llmInstance := openai.CreateGPtInstance(openai.Config{
		OrgId:     os.Getenv("GPT_ORG_ID"),
		ApiKey:    os.Getenv("GPT_API_KEY"),
		ProjectId: "ideography-llm", // for now hardcoding the project id
	})
	return llmInstance
}
