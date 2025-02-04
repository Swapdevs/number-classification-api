package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"

)

type NumberResponse struct {
	Number    int      `json:"number"`
	IsPrime   bool     `json:"is_prime"`
	IsPerfect bool     `json:"is_perfect"`
	Properties []string `json:"properties"`
	DigitSum  int      `json:"digit_sum"`
	FunFact   string   `json:"fun_fact"`
}

type ErrorResponse struct {
	Number string `json:"number"`
	Error  bool   `json:"error"`
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isPerfect(n int) bool {
	if n < 2 {
		return false
	}
	sum := 1
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			sum += i
			if i != n/i {
				sum += n / i
			}
		}
	}
	return sum == n
}

func isArmstrong(n int) bool {
	digits := strconv.Itoa(n)
	length := len(digits)
	sum := 0
	for _, digit := range digits {
		d, _ := strconv.Atoi(string(digit))
		sum += int(math.Pow(float64(d), float64(length)))
	}
	return sum == n
}

func digitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func classifyNumber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	numberStr := r.URL.Query().Get("number")
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		errorResponse := ErrorResponse{Number: numberStr, Error: true}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	properties := []string{}
	if isArmstrong(number) {
		properties = append(properties, "armstrong")
	}
	if number%2 == 0 {
		properties = append(properties, "even")
	} else {
		properties = append(properties, "odd")
	}

	funFact := fmt.Sprintf("%d is an interesting number.", number)
	if isArmstrong(number) {
		funFact = fmt.Sprintf("%d is an Armstrong number because the sum of its own digits each raised to the power of the number of digits equals itself.", number)
	}

	response := NumberResponse{
		Number:    number,
		IsPrime:   isPrime(number),
		IsPerfect: isPerfect(number),
		Properties: properties,
		DigitSum:  digitSum(number),
		FunFact:   funFact,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/api/classify-number", classifyNumber)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
