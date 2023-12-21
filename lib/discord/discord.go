package discord

import(
  "encoding/base64"
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "bytes"
)

func Gathering(token string) {
  UserData(token)
  PaymentMethod(token)
  Friend(token)
  Channel(token)
}

func Token(id string) {
  fmt.Println("\nFirst token part: ", base64.StdEncoding.EncodeToString([]byte(id))+".xxxxxx.xxxxxxxxxxxxxxxxxxxxxxxxxx\n")
}

func fetchData(url, token string) ([]byte, error) {
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11")
	req.Header.Add("Authorization", token)
	req.Header.Add("Cookie", "")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func prettyPrint(body []byte) (string, error) {
	// Marshal the response body into an indented JSON string
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, body, "", "  ")
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return "", err
	}

	// Return the prettified JSON as a string
	return prettyJSON.String(), nil
}

func PaymentMethod(token string) {
  url := "https://discordapp.com/api/v6/users/@me/billing/payment-sources"
  body, err := fetchData(url, token)
	if err != nil {
		fmt.Println(err)
		return
	}
  fmt.Println(prettyPrint(body))
}

func Friend(token string) {
  url := "https://discordapp.com/api/v6/users/@me/relationships"
  body, err := fetchData(url, token)
	if err != nil {
		fmt.Println(err)
		return
	}
  fmt.Println(prettyPrint(body))
}

func UserData(token string) {
  url := "https://discordapp.com/api/v6/users/@me"
  body, err := fetchData(url, token)
	if err != nil {
		fmt.Println(err)
		return
	}
  fmt.Println(prettyPrint(body))
}

func Channel(token string) {
  url := "https://discordapp.com/api/v6/users/@me/channels"
  body, err := fetchData(url, token)
	if err != nil {
		fmt.Println(err)
		return
	}
  fmt.Println(prettyPrint(body))
}