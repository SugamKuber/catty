package handlers

import (
	"context"
	"fmt"
	"os"
	"server/internal/services"
	"server/internal/types"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"net/http"
)

func Generate(c *gin.Context, db *types.Database) {
	employeeID, exists := c.Get("employee_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found in context"})
		return
	}

	var user struct {
		EmployeeID    string    `bson:"employee_id"`
		Email         string    `bson:"email"`
		InspectorName string    `bson:"inspector_name"`
	}
	userCollection := db.MongoClient.Database("catty").Collection("users")

	err := userCollection.FindOne(context.Background(), bson.M{"employee_id": employeeID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user data"})
		}
		return
	}

	err = services.GeneratePDF(user.InspectorName, user.Email)
	if err != nil {
		fmt.Printf("Error generating PDF: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("PDF generated successfully!")

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "mail sent"})
}
