package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

// Person struct to demonstrate methods and struct usage
type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

// Method to get the full name of a person
func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

// Function to handle HTTP GET requests
func handlerGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

// Function to handle HTTP POST requests
func handlerPost(w http.ResponseWriter, r *http.Request) {
	var p Person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := encodeToJSON(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Received Person: %s\n", response)
}

// Function to read a file and return its contents
func readFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Function to write content to a file
func writeFile(filename, content string) error {
	return ioutil.WriteFile(filename, []byte(content), 0644)
}

// Goroutine function to print numbers
func printNumbers(label string, wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		msg := fmt.Sprintf("%s: %d", label, i)
		fmt.Println(msg)
		ch <- msg
		time.Sleep(1 * time.Second)
	}
}

// Function to encode a Person struct to JSON
func encodeToJSON(p Person) (string, error) {
	jsonData, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

// Function to decode JSON to a Person struct
func decodeFromJSON(data string) (Person, error) {
	var p Person
	err := json.Unmarshal([]byte(data), &p)
	if err != nil {
		return Person{}, err
	}
	return p, nil
}

// Function to make a simple HTTP GET request
func makeHTTPRequest(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// Function to parse time from a string
func parseTime(timeStr string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

// Function to reverse a string
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Function to check if a string is a palindrome
func isPalindrome(s string) bool {
	return s == reverseString(s)
}

// Function to concatenate two strings
func concatenateStrings(s1, s2 string) string {
	return s1 + s2
}

// Coroutine function to handle string operations
func handleStringOperations(s string, wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()
	reversed := reverseString(s)
	palindrome := isPalindrome(s)
	result := fmt.Sprintf("Original: %s, Reversed: %s, IsPalindrome: %v", s, reversed, palindrome)
	ch <- result
}

// Main function
func main() {
	// Basic I/O operations
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Printf("Hello, %s!\n", name)

	// Working with structs and methods
	person := Person{FirstName: "John", LastName: "Doe", Age: 30}
	fmt.Printf("Person: %s, Age: %d\n", person.FullName(), person.Age)

	// Encoding to JSON
	jsonString, err := encodeToJSON(person)
	if err != nil {
		log.Fatalf("Error encoding to JSON: %v", err)
	}
	fmt.Println("JSON:", jsonString)

	// Decoding from JSON
	decodedPerson, err := decodeFromJSON(jsonString)
	if err != nil {
		log.Fatalf("Error decoding from JSON: %v", err)
	}
	fmt.Printf("Decoded Person: %s, Age: %d\n", decodedPerson.FullName(), decodedPerson.Age)

	// Reading from a file
	content, err := readFile("example.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	fmt.Println("File content:", content)

	// Writing to a file
	err = writeFile("output.txt", "This is a test content")
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}
	fmt.Println("Content written to output.txt")

	// Concurrency with goroutines and channels
	var wg sync.WaitGroup
	ch := make(chan string)
	wg.Add(2)
	go printNumbers("Goroutine 1", &wg, ch)
	go printNumbers("Goroutine 2", &wg, ch)
	go func() {
		for msg := range ch {
			fmt.Println("Channel received:", msg)
		}
	}()
	wg.Wait()
	close(ch)

	// String operations with goroutines
	stringCh := make(chan string)
	wg.Add(2)
	go handleStringOperations("racecar", &wg, stringCh)
	go handleStringOperations("hello", &wg, stringCh)
	go func() {
		for result := range stringCh {
			fmt.Println("String operation result:", result)
		}
	}()
	wg.Wait()
	close(stringCh)

	// HTTP GET request
	response, err := makeHTTPRequest("http://www.google.com")
	if err != nil {
		log.Fatalf("Error making HTTP request: %v", err)
	}
	fmt.Println("HTTP Response:", response)

	// Time and date operations
	now := time.Now()
	fmt.Println("Current Time:", now.Format(time.RFC1123))

	// Parsing time from a string
	parsedTime, err := parseTime("2024-06-06 12:00:00")
	if err != nil {
		log.Fatalf("Error parsing time: %v", err)
	}
	fmt.Println("Parsed Time:", parsedTime)

	// Start a simple HTTP server
	http.HandleFunc("/", handlerGet)
	http.HandleFunc("/post", handlerPost)
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
	fmt.Println("Server started at http://localhost:8080")

	// Keep the main function running to keep the server running
	select {}
}
