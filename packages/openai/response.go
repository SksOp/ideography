package openai

type Usage struct {
	Prompt_Token     int `json:"prompt_token"`
	Completion_Token int `json:"Completion_Token"`
	Total_Token      int `json:"total_token"`
}

type Choices struct {
	Message Message `json:"message"`
	//  Logprobs nil `json:"logprobs"`
	Finish_Reason string `json:"finish_reason"`
	Index         int    `json:"index"`
}

type GptResponse struct {
	Id      string    `json:"id"`
	Object  string    `json:"object"`
	Created int       `json:"created"`
	Model   string    `json:"Model"`
	Usage   Usage     `json:"usage"`
	Choices []Choices `json:"choices"`
}
