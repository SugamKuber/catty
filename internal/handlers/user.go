package handlers

import (
	"context"
	"net/http"
	"regexp"
	"server/internal/types"
	"server/utils"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func isNumeric(s string) bool {
	match, _ := regexp.MatchString(`^\d+$`, s)
	return match
}

func Signup(c *gin.Context, db *types.Database) {
	var user struct {
		Password      string `json:"password"`
		EmployeeID    string `json:"employee_id"`
		InspectorName string `json:"inspector_name"`
		Email         string `json:"email"`
	}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if !isNumeric(user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be numeric only"})
		return
	}
	userCollection := db.MongoClient.Database("catty").Collection("users")

	if userCollection == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Users not initialized"})
		return
	}

	var existingUser struct {
		ID            primitive.ObjectID `bson:"_id,omitempty"`
		Password      string             `bson:"password"`
		EmployeeID    string             `bson:"employee_id"`
		InspectorName string             `bson:"inspector_name"`
		Email         string             `bson:"email"`
	}

	err := userCollection.FindOne(context.Background(), bson.M{"employee_id": user.EmployeeID}).Decode(&existingUser)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}
	if err != mongo.ErrNoDocuments {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking user: " + err.Error()})
		return
	}

	_, err = userCollection.InsertOne(context.Background(), bson.M{
		"password":       user.Password,
		"employee_id":    user.EmployeeID,
		"inspector_name": user.InspectorName,
		"email":          user.Email,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Signup successful"})
}

func Login(c *gin.Context, db *types.Database) {
	var user struct {
		EmployeeID string `json:"employee_id"`
		Password   string `json:"password"`
	}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if !isNumeric(user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be numeric only"})
		return
	}

	var existingUser struct {
		ID       primitive.ObjectID `bson:"_id,omitempty"`
		Password string             `bson:"password"`
	}
	userCollection := db.MongoClient.Database("catty").Collection("users")

	err := userCollection.FindOne(context.Background(), bson.M{"employee_id": user.EmployeeID}).Decode(&existingUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid employee_id or password"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking user"})
		}
		return
	}

	if existingUser.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid employee_id or password"})
		return
	}

	token, err := utils.GenerateToken(user.EmployeeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}

func GetUserData(c *gin.Context, db *types.Database) {
	employeeID, exists := c.Get("employee_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found in context"})
		return
	}

	var user struct {
		EmployeeID    string    `bson:"employee_id"`
		InspectorName string    `bson:"inspector_name"`
		CreatedAt     time.Time `bson:"created_at"`
		Email         string    `bson:"email"`
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

	c.JSON(http.StatusOK, gin.H{
		"employee_id":    user.EmployeeID,
		"inspector_name": user.InspectorName,
		"created_at":     user.CreatedAt,
		"email":          user.Email,
	})
}
