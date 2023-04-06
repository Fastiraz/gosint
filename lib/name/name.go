package name

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var printf = fmt.Printf

type Website struct {
	ErrorType         string `json:"errorType"`
	Url               string `json:"url"`
	UrlMain           string `json:"urlMain"`
	nameClaimed   string `json:"name_claimed"`
}

func uname(name string) {
	// Read the JSON file
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	// Unmarshal the JSON data into a map
	websites := make(map[string]Website)
	err = json.Unmarshal(data, &websites)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		os.Exit(1)
	}

	// Prompt the user for a name
	//var name string
	//fmt.Print("Enter a name: ")
	//fmt.Scanln(&name)

	// Replace the "{}" in the URLs with the name and make an HTTP GET request to each URL
	for name, website := range websites {
		url := strings.Replace(website.Url, "{}", name, 1)
		response, err := http.Get(url)
		if err != nil {
			fmt.Printf("\x1b[31m%s: Error making request: %v\x1b[0m\n", name, err)
			continue
		}
		defer response.Body.Close()

		statusCode := response.StatusCode
		if statusCode >= 200 && statusCode < 300 {
			fmt.Printf("\x1b[32m%s: %s\x1b[0m\n", name, response.Status)
		} else if statusCode >= 400 && statusCode < 500 {
			fmt.Printf("\x1b[31m%s: %s\x1b[0m\n", name, response.Status)
		} else {
			fmt.Printf("\x1b[33m%s: %s\x1b[0m\n", name, response.Status)
		}
	}

	fmt.Println("Requests completed successfully.")
}

func Name(name string) {
  // Read the JSON file
  data, err := ioutil.ReadFile("./lib/name/data.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	// Unmarshal the JSON data into a map
	websites := make(map[string]Website)
	err = json.Unmarshal(data, &websites)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		os.Exit(1)
	}

	// Prompt the user for a name
	//var name string
	//fmt.Print("Enter a name: ")
	//fmt.Scanln(&name)

	// Replace the "{}" in the URLs with the name and make an HTTP GET request to each URL
	for name, website := range websites {
		url := strings.Replace(website.Url, "{}", name, 1)
		response, err := http.Get(url)
		if err != nil {
			fmt.Printf("\x1b[31m%s: Error making request: %v\x1b[0m\n", name, err)
			continue
		}
		defer response.Body.Close()

		statusCode := response.StatusCode
		if statusCode >= 200 && statusCode < 300 {
			fmt.Printf("\x1b[32m%s: %s\x1b[0m\n", name, response.Status)
		} else if statusCode >= 400 && statusCode < 500 {
			fmt.Printf("\x1b[31m%s: %s\x1b[0m\n", name, response.Status)
		} else {
			fmt.Printf("\x1b[33m%s: %s\x1b[0m\n", name, response.Status)
		}
	}

	fmt.Println("Requests completed successfully.")
}
