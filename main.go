package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	expirationDate = "2025-12-31"    // Expiration date for the license
	dateFormat     = "20060102_1504" // Format for the log file timestamp
)

// banner displays the program banner and waits for 5 seconds
func banner() {
	front := `
     _                                                        
    /   ._   _    _.  _|_   _    _|    |_                     
    \_  |   (/_  (_|   |_  (/_  (_|    |_)  \/                
                                            /                 
          ___                      __                ___  ___ 
    |\ |   |   |\ |    |   /\     (_    /\   \    /   |    |  
    | \|  _|_  | \|  \_|  /--\    __)  /--\   \/\/   _|_   |  
	`
	fmt.Println(front)
	time.Sleep(5 * time.Second)
}

// currentDateTime returns the current date and time in the specified format
func currentDateTime() string {
	return time.Now().Format(dateFormat)
}

// createLog creates or appends to a log file with the given content
func createLog(content, y string) {
	// Ensure the "result" directory exists
	if err := os.MkdirAll("./result", os.ModePerm); err != nil {
		fmt.Println("Error creating result directory:", err)
		return
	}

	// Create the log file name
	now := currentDateTime()
	logFileName := fmt.Sprintf("./result/%s_%s.txt", y, now)

	// Open the log file in append mode
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer logFile.Close()

	// Write the content to the log file
	if _, err := logFile.WriteString(content + "\n"); err != nil {
		fmt.Println("Error writing to log file:", err)
	}
}

// checkContent walks through the directory and checks for files with the specified extension
func checkContent(startPath, val2 string) {
	// Read payloads from the payload.txt file
	payloads, err := ioutil.ReadFile("payloads.txt")
	if err != nil {
		fmt.Println("Error reading payloads.txt:", err)
		return
	}

	// Split payloads into lines
	payloadLines := strings.Split(string(payloads), "\n")

	// Walk through the directory
	err = filepath.Walk(startPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error accessing path:", path, err)
			return nil
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Check if the file has the specified extension
		if strings.HasSuffix(path, "."+val2) {
			// Read the file content
			fileContent, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Println("Error reading file:", path, err)
				return nil
			}

			// Check for payloads in the file content
			var foundPayloads []string
			for _, payload := range payloadLines {
				if strings.Contains(string(fileContent), payload) && payload != "" {
					foundPayloads = append(foundPayloads, payload)
				}
			}

			// If payloads are found, log the result with commas
			if len(foundPayloads) >= 1 {
				fmt.Println(len(foundPayloads))
				// Join the foundPayloads array with commas
				payloadsString := strings.Join(foundPayloads, ", ")
				result := fmt.Sprintf("%s ==> [%s]", path, payloadsString)
				createLog(result, val2)
				fmt.Println(result)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking through directory:", err)
	}
}

func main() {
	// Check command-line arguments
	if len(os.Args) < 3 {
		fmt.Println("Usage: ./scan </path/of/directory> <extention>")
		return
	}

	val1 := os.Args[1] // Start path
	val2 := os.Args[2] // File extension to check

	// Check if the license has expired
	currentDate := time.Now().Format("2006-01-02")
	if currentDate >= expirationDate {
		fmt.Println("Renew License Soon")
		return
	}

	// Display the banner and start checking content
	banner()
	checkContent(val1, val2)
}
