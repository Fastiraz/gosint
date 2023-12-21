package discord

import(
  "encoding/base64"
  "fmt"
  "net/http"
  "io/ioutil"
)

func Token(id string) {
  encodedID := base64.StdEncoding.EncodeToString([]byte(id))
  fmt.Println("\nFirst token part: ", encodedID+".xxxxxx.xxxxxxxxxxxxxxxxxxxxxxxxxx\n")
}

func Friend(token string) {
  url := "https://discordapp.com/api/v6/users/@me/relationships"
  method := "GET"

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, nil)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Content-Type", "application/json")
  req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11")
  req.Header.Add("Authorization", token)
  req.Header.Add("Cookie", "__dcfduid=3fe768cc9f4111ee97fd0a41db119d69; __sdcfduid=3fe768cc9f4111ee97fd0a41db119d69290da082faf8c4657a331f18c9e367c4a6035b76c5be5c405fee27a3e169b8ca")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}
