// package main

// import (
// 	"bufio"
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

// type User struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// 	Role     string `json:"role"`
// }

// func main() {
// 	// Open the input file
// 	file, err := os.Open("temp.json")
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return
// 	}
// 	defer file.Close()

// 	var users []User
// 	scanner := bufio.NewScanner(file)

// 	for scanner.Scan() {
// 		var user User
// 		line := scanner.Text()
// 		if err := json.Unmarshal([]byte(line), &user); err != nil {
// 			fmt.Println("Skipping invalid line:", line)
// 			continue
// 		}
// 		users = append(users, user)
// 	}

// 	// Optional: check for scanner errors
// 	if err := scanner.Err(); err != nil {
// 		fmt.Println("Error reading file:", err)
// 		return
// 	}

// 	// Create the output file
// 	outFile, err := os.Create("temp.json")
// 	if err != nil {
// 		fmt.Println("Error creating output file:", err)
// 		return
// 	}
// 	defer outFile.Close()

// 	// Encode the slice of users as a JSON array
// 	encoder := json.NewEncoder(outFile)
// 	encoder.SetIndent("", "  ") // pretty-print
// 	if err := encoder.Encode(users); err != nil {
// 		fmt.Println("Error writing JSON:", err)
// 	}
// }
