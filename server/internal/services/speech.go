package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func ConvertSpeechToText(file io.Reader) (string, error) {
	var formData bytes.Buffer
	writer := multipart.NewWriter(&formData)

	part, err := writer.CreateFormFile("file", "audio.mp3")
	if err != nil {
		return "", err
	}

	if _, err := io.Copy(part, file); err != nil {
		return "", err
	}

	err = writer.WriteField("model", "whisper-1")
	if err != nil {
		return "", err
	}

	writer.Close()
	APIURL := "https://api.openai.com/v1/audio/transcriptions"
	req, err := http.NewRequest("POST", APIURL, &formData)
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "Bearer "+os.Getenv("OPENAPI_TOKEN"))
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	if text, ok := result["text"].(string); ok {
		return text, nil
	}

	return "", fmt.Errorf("unexpected response format")
}
