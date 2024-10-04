package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"server/internal/types"
)

type OpenAIRequest struct {
	Model     string    `json:"model"`
	Messages  []Message `json:"messages"`
	MaxTokens int       `json:"max_tokens"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

func EvaluateInspectionData(inputData types.Content) (string, error) {
	url := "https://api.openai.com/v1/chat/completions"

	requestBody := OpenAIRequest{
		Model: "gpt-4",
		Messages: []Message{
			{
				Role:    "system",
				Content: "1. You are my Caterpillar vechical (catty for short) inspector for wheels. 2. give me only one 1 line answers or max 2 line answers. 3. Give me the status of the tires for PSI & condition (tire wear), the user will give the required data for vechicals like truck 4.Dont give me more than 3 lines answer 5. dont greet & thank user, give the data direclty 6. the tire pressure should be between 70 to 80 PSI this is the good amount for the CAT vehical 7. Tire condition can be good or need replacement, Take this input when the inspec	tor says about the tire wear (tire grip amount)",
			},
			{
				Role:    "user",
				Content: fmt.Sprintf("I am an inspector at Caterpillar, an automotive and machine company. Please evaluate the following component inspection for tire data PSI & wear, give response in 1 line only: %s", inputData),
			},
		},
		MaxTokens: 100,
	}

	requestBodyJson, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshaling request body")
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyJson))
	if err != nil {
		fmt.Println("Error creating HTTP request")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAPI_TOKEN"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request")
		return "", err

	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body")
		return "", err

	}

	var openAIResponse OpenAIResponse
	err = json.Unmarshal(body, &openAIResponse)
	if err != nil {
		fmt.Println("Error unmarshaling response body")
		return "", err
	}

	fmt.Println("Response from GPT-4:", openAIResponse.Choices[0].Message.Content)
	return openAIResponse.Choices[0].Message.Content, nil
}
