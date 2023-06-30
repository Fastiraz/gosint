package discord

import(
  "encoding/base64"
  "fmt"
)

func Token(id string) {
  encodedID := base64.StdEncoding.EncodeToString([]byte(id))
  fmt.Println("\nFirst token part: ", encodedID+".xxxxxx.xxxxxxxxxxxxxxxxxxxxxxxxxx\n")
}
