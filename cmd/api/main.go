package main

import (
	"fmt"
	"ideography/internal/server"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

//    gptInstance := gpt.CreateGPtInstance(gpt.Config{
//    	OrgId: os.Getenv("GPT_ORG_ID"),
//    	ApiKey: os.Getenv("GPT_API_KEY"),
//    	ProjectId: "sdf",
//    })
//    gptInstance.Query("hi")
}
