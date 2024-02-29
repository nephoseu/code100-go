package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func helloworld() string {
	return "Hello World!!"
}

func main() {

	email := "EMAIL-OR-USERNAME-HERE" // replace with email or username given to you
	password := "PASSWORD-HERE"       // replace with password given to you

	// Step 1: Login to the API endpoint
	credentials, err := json.Marshal(map[string]string{
		"email":    email,
		"password": password,
	})
	if err != nil {
		fmt.Println("Error: ", err)
	}

	response, err := http.Post("https://challenger.code100.dev/login", "application/json", bytes.NewReader(credentials))
	if err != nil {
		fmt.Println("Error Login: ", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		fmt.Println("Error: ", response.Status)
		return
	}

	var result map[string]interface{}
	json.NewDecoder(response.Body).Decode(&result)
	token := result["token"].(string)

	// Step 2: Call Authenticated Endpoint /testauthroute
	req, err := http.NewRequest("GET", "https://challenger.code100.dev/testauthroute", nil)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err = client.Do(req)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	sb := string(body)
	fmt.Println("Response: ", sb)

	// Step 3: Get the puzzle
	req, err = http.NewRequest("GET", "https://challenger.code100.dev/getpuzzle", nil)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	client = &http.Client{}
	response, err = client.Do(req)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer response.Body.Close()

	body, err = io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	sb = string(body)
	fmt.Println("Puzzle: ", sb)

	// Step 4: Solve the puzzle

	////////////////////////////
	////// YOUR CODE HERE //////
	////////////////////////////

	answer, err := json.Marshal(map[string]string{
		"answer": "Answer in the required format",
	})
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// Step 5: Submit the solution
	req, err = http.NewRequest("POST", "https://challenger.code100.dev/postanswer", bytes.NewReader(answer))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	client = &http.Client{}
	response, err = client.Do(req)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer response.Body.Close()

	body, err = io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	sb = string(body)
	fmt.Println("Response: ", sb)

}
