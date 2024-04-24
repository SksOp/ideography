package main

import (
	"fmt"
//	"ideography/internal/server"
"ideography/packages/gpt"
"os"
_ "github.com/joho/godotenv/autoload"
)

func main() {

//	server := server.NewServer()
//
//	err := server.ListenAndServe()
//	if err != nil {
//		panic(fmt.Sprintf("cannot start server: %s", err))
//	}
	api := os.Getenv("GPT_API_KEY")
	fmt.Println(api)
   gptInstance := gpt.CreateGPtInstance(gpt.Config{
   	OrgId: os.Getenv("GPT_ORG_ID"),
   	ApiKey: os.Getenv("GPT_API_KEY"),
   	ProjectId: "sdf",
   })
   gptInstance.Query("hi")
}
