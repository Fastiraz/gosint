package github

import (
        "encoding/json"
        "fmt"
        "io/ioutil"
        "net/http"
)

func Email(username string) {
        url := fmt.Sprintf("https://api.github.com/users/%s/events/public", username)

        response, err := http.Get(url)
        if err != nil {
                fmt.Println("\nError making the request:", err, "\n")
                return
        }
        defer response.Body.Close()

        body, err := ioutil.ReadAll(response.Body)
        if err != nil {
                fmt.Println("\nError reading the response body:", err, "\n")
                return
        }

        var events []map[string]interface{}
        err = json.Unmarshal(body, &events)
        if err != nil {
                fmt.Println("\nError unmarshaling the response:", err, "\n")
                return
        }

        if len(events) > 0 {
                payload := events[0]["payload"].(map[string]interface{})
                commits := payload["commits"].([]interface{})
                if len(commits) > 0 {
                        author := commits[0].(map[string]interface{})["author"].(map[string]interface{})
                        email := author["email"].(string)
                        fmt.Println("\nEmail:", email, "\n")
                        return
                }
        }

        fmt.Println("\nNo data or email found.\n")
}
