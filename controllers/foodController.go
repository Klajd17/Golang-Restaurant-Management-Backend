package controllers

import "github.com/gin-gonic/gin"

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Get Foods",
		})
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Get Food",
		})
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Create Food",
		})
	}
}

func round(num float64) int {

}

func toFixed(num float64, precision int) float64 {

}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Update Food",
		})
	}
}

func DeleteFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Delete Food",
		})
	}
}
