package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Function to check if a number is prime
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Function to check if a number is perfect
func isPerfect(n int) bool {
	if n <= 0 { // Ensure only positive numbers are checked
		return false
	}
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum == n
}

// Function to check if a number is an Armstrong number
func isArmstrong(n int) bool {
	sum := 0
	temp := n
	digits := len(strconv.Itoa(n))

	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		temp /= 10
	}
	return sum == n
}

// Function to get the sum of digits of a number
func digitSum(n float64) int {
	sum := 0
	for _, digit := range strconv.FormatFloat(math.Abs(n), 'f', -1, 64) {
		if digit >= '0' && digit <= '9' {
			sum += int(digit - '0')
		}
	}
	return sum
}

// Function to classify the number
func classifyNumber(c *gin.Context) {
	numStr := c.Query("number")

	// Set JSON Content-Type explicitly for all responses
	c.Writer.Header().Set("Content-Type", "application/json")

	// Convert input to float (to handle both integers and floats)
	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		// If input is not a number, return 400
		c.JSON(http.StatusBadRequest, gin.H{
			"number": numStr,
			"error":  true,
		})
		return
	}

	// Determine properties
	properties := []string{}
	if int(math.Abs(num))%2 == 0 {
		properties = append(properties, "even")
	} else {
		properties = append(properties, "odd")
	}
	if isArmstrong(int(num)) {
		properties = append(properties, "armstrong")
	}

	// Fun fact (Replace with real API later)
	funFact := fmt.Sprintf("%.2f is an interesting number!", num)

	// Send JSON response
	c.JSON(http.StatusOK, gin.H{
		"number":     num,
		"is_prime":   isPrime(int(num)),
		"is_perfect": isPerfect(int(num)), // Now correctly handles 0
		"properties": properties,
		"digit_sum":  digitSum(num),
		"fun_fact":   funFact,
	})
}

func main() {
	r := gin.Default()

	// CORS Middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Ensure all responses include "Content-Type: application/json"
		c.Writer.Header().Set("Content-Type", "application/json")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	// API Route
	r.GET("/api/classify-number", classifyNumber)

	// Start server on port 8080
	r.Run(":8080")
}

