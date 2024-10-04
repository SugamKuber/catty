package handlers

import (
	"fmt"
	"io"

	"github.com/gin-gonic/gin"

	"net/http"
	"server/internal/services"
	"server/internal/types"
)

func Speech(c *gin.Context) {
	fmt.Println("this is input")
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Please give the correct input"})
		return
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Could not read file"})
		return
	}

	fileType := http.DetectContentType(buffer)
	file.Seek(0, io.SeekStart)

	fmt.Println("File Name:", fileHeader.Filename)
	fmt.Println("File Size:", fileHeader.Size)
	fmt.Println("File Type:", fileType)
	fmt.Println("File Header:", fileHeader.Header)

	// vehicle := c.PostForm("vehicle")
	// component := c.PostForm("component")
	vehicle := "truck loader"
	component := "front right tire & right rear tire"

	if vehicle == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Vehicle is required"})
		return
	}
	if component == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Component is required"})
		return
	}

	transcription, err := services.ConvertSpeechToText(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to convert speech to text"})
		return
	}

	content := types.Content{
		Message:   transcription,
		Component: component,
		Vehicle:   vehicle,
	}

	fmt.Println(content)

	response, err := services.EvaluateInspectionData(content)
	if err != nil {
		fmt.Println("error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to convert get the llm response."})
		return
	}
	fmt.Println(response)

	audioResponse, err := services.ConvertTextToSpeech(response)
	if err != nil {
		fmt.Println("error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to convert text to speech."})
		return
	}
	defer audioResponse.Body.Close()

	audioData, err := io.ReadAll(audioResponse.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to read audio data."})
		return
	}
	fmt.Println("file generated")
	fileName := "reply.wav"
	c.Header("Content-Type", "audio/wav")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Data(http.StatusOK, "audio/wav", audioData)
	
}
