package email

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
  	"strings"
)

func spotify(email string, client *http.Client, out chan<- map[string]interface{}) {
	name := "spotify"
	domain := "spotify.com"
	method := "register"
	frequentRateLimit := true

	headers := map[string]string{
		"User-Agent":      getRandomUserAgent(),
		"Accept":          "application/json, text/plain, */*",
		"Accept-Language": "en-US,en;q=0.5",
		"DNT":             "1",
		"Connection":      "keep-alive",
	}

	params := url.Values{
		"validate": {"1"},
		"email":    {email},
	}

	req, err := http.NewRequest("GET", "https://spclient.wg.spotify.com/signup/public/v1/account", nil)
	if err != nil {
		out <- createResultMap(name, domain, method, frequentRateLimit, true, nil, nil, nil, nil)
		return
	}

	req.URL.RawQuery = params.Encode()

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		out <- createResultMap(name, domain, method, frequentRateLimit, false, nil, nil, nil, nil)
		return
	}
	defer resp.Body.Close()

	// TODO: Perform JSON parsing and check the value of "status" field in the response

	out <- createResultMap(name, domain, method, frequentRateLimit, false, nil, nil, nil, nil)
}

func laposte(email string, client *http.Client, out chan<- map[string]interface{}) {
	name := "laposte"
	domain := "laposte.fr"
	method := "register"
	frequentRateLimit := false

	headers := map[string]string{
		"Origin":          "https://www.laposte.fr",
		"User-Agent":      getRandomUserAgent(),
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		"Referer":         "https://www.laposte.fr/authentification",
		"Accept-Language": "fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7",
	}

	data := url.Values{
		"email":       {email},
		"customerId":  {""},
		"tunnelSteps": {""},
	}

	req, err := http.NewRequest("POST", "https://www.laposte.fr/authentification", strings.NewReader(data.Encode()))
	if err != nil {
		out <- createResultMap(name, domain, method, frequentRateLimit, true, false, nil, nil, nil)
		return
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		out <- createResultMap(name, domain, method, frequentRateLimit, false, false, nil, nil, nil)
		return
	}
	defer resp.Body.Close()

	exists := false
	// TODO: Perform HTML parsing to check if the email exists

	out <- createResultMap(name, domain, method, frequentRateLimit, false, exists, nil, nil, nil)
}

func getRandomUserAgent() string {
	ua := []string{
		"Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2227.1 Safari/537.36",
	}

	return ua[rand.Intn(len(ua))]
}

func createResultMap(name, domain, method string, frequentRateLimit bool, rateLimit, exists, emailRecovery, phoneNumber, others interface{}) map[string]interface{} {
	return map[string]interface{}{
		"name":                name,
		"domain":              domain,
		"method":              method,
		"frequent_rate_limit": frequentRateLimit,
		"rateLimit":           rateLimit,
		"exists":              exists,
		"emailrecovery":       emailRecovery,
		"phoneNumber":         phoneNumber,
		"others":              others,
	}
}

/*
func main() {
	client := &http.Client{}
	out := make(chan map[string]interface{})

	emailAddress := "test@example.com"
	
  fmt.Println("\nChecking for spotify: {}", emailAddress)
	go spotify(emailAddress, client, out)
  result := <-out
	fmt.Println(result)
  
  fmt.Println("\nChecking for laposte: {}", emailAddress)
  go laposte(emailAddress, client, out)
	result2 := <-out
	fmt.Println(result2)
}
*/

func Email(emailAddress string) {
	client := &http.Client{}
	out := make(chan map[string]interface{})

	//emailAddress := "test@example.com"

	fmt.Printf("\nChecking for spotify: %s\n", emailAddress)
	go spotify(emailAddress, client, out)
	result := <-out
	fmt.Printf("Name: %s\nDomain: %s\nMethod: %s\nFrequent Rate Limit: %t\nRate Limit: %v\nExists: %v\nEmail Recovery: %v\nPhone Number: %v\nOthers: %v\n\n", 
		result["name"], result["domain"], result["method"], result["frequent_rate_limit"], result["rateLimit"], 
		result["exists"], result["emailrecovery"], result["phoneNumber"], result["others"])

	fmt.Printf("\nChecking for laposte: %s\n", emailAddress)
	go laposte(emailAddress, client, out)
	result2 := <-out
	fmt.Printf("Name: %s\nDomain: %s\nMethod: %s\nFrequent Rate Limit: %t\nRate Limit: %v\nExists: %v\nEmail Recovery: %v\nPhone Number: %v\nOthers: %v\n\n", 
		result2["name"], result2["domain"], result2["method"], result2["frequent_rate_limit"], result2["rateLimit"], 
		result2["exists"], result2["emailrecovery"], result2["phoneNumber"], result2["others"])
}
