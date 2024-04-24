package gpt

import (
 "net/http" 
"fmt"
"bytes"
"io/ioutil"
"encoding/json"
)

type Config struct {
 OrgId     string
 ProjectId string
 ApiKey    string
}

func CreateGPtInstance (c Config) *GPT {
 return &GPT{c,"gpt-3.5-turbo", "https://api.openai.com/v1/chat/completions",0.7}
}

type GPT struct {
 config Config
 model string
 url string
 temperature float64
}



func (g *GPT) Query(q string)  {
  
  message:= CreateUserMessageObject(q)
  messages:= make([]Message,0)
  messages=append(messages,message)
  
  body:=g.CreateQueryBody(messages)
  
  client:= &http.Client{}
  
  req,_:= http.NewRequest("POST",g.url,git a)
  authheader:= fmt.Sprintf("Bearer %v",g.config.ApiKey)

  req.Header.Set("Authorization",authheader)
  req.Header.Set("Content-Type","application/json")
  resp,_:= client.Do(req)

  defer resp.Body.Close() 
//  fmt.Println(resp.Body)
  
  b,_ := ioutil.ReadAll(resp.Body)
  
  
  
  var response GptResponse;
  
  if err := json.Unmarshal(b, &response); err != nil {
    fmt.Println(err)
    return 
  }
  
  fmt.Println(response)
  
  
}