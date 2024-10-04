package types

type Content struct {
	Message   string `json:"message" binding:"required"`
	Component string `json:"component" binding:"required"`
	Vehicle   string `json:"vehicle" binding:"required"`
}