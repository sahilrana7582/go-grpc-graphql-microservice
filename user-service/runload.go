package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	file, err := os.Open("output.json")
	if err != nil {
		fmt.Println("❌ Error opening temp.json:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		count++
		fmt.Printf("🚀 Running ghz test for user #%d\n", count)

		// Save the line to a temporary file
		err := os.WriteFile("line.json", []byte(scanner.Text()), 0644)
		if err != nil {
			fmt.Println("❌ Error writing line.json:", err)
			return
		}

		// Run ghz on that line
		cmd := exec.Command("ghz",
			"--insecure",
			"--proto", "./proto/user.proto",
			"--call", "user.UserService.CreateUser",
			"-d", "temp.json",
			"-c", "1",
			"-n", "1",
			"localhost:50051",
		)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Printf("❌ ghz failed for user #%d: %v\n", count, err)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("❌ Error reading temp.json:", err)
	} else {
		fmt.Printf("✅ Done testing %d users.\n", count)
	}

}
