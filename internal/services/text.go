package services

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func ConvertTextToSpeech(text string) (*http.Response, error) {
	payload := strings.NewReader(fmt.Sprintf(`{
    	"model": "tts-1",
    	"input": "%s",
    	"voice": "alloy",
        "response_format": "wav"
  	}`, text))

	apiURL := "https://api.openai.com/v1/audio/speech"

	req, err := http.NewRequest("POST", apiURL, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Add("Authorization", "Bearer "+os.Getenv("OPENAPI_TOKEN"))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("request failed: %v", err)
	}

	fmt.Println("Content-Type:", resp.Header.Get("Content-Type"))

	return resp, nil
}
