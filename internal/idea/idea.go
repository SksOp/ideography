package idea

import (
	LLM "ideography/internal/llm"
)

func HandleNewIdea(q string) (string, error) {

	system := `
		You are a bot that take the raw idea from the user and 
		process it to make it more meaningful and actionable.

		You can use the following steps to process the idea:
		1. Understand the context of the idea
		2. Break the idea into smaller parts
		3. Analyze the parts and find the core problem
		4. Find the solution to the core problem
		5. Combine the solution to get the final idea
		6. Think about the risks and mitigations
		7. Write the final idea in a structured format

		Don't start like its a nice idea, but directly start with 
		# Introduction
		......

		Here is the User Idea:
	`
	llm := LLM.NewLLM()
	systemMessage := llm.CreateSystemMessageObject(system)
	userMessage := llm.CreateUserMessageObject(q)

	query := []LLM.Message{systemMessage, userMessage}

	response, err := llm.Query(query)

	if err != nil {
		return "", err
	}

	return response.Content, nil

}
