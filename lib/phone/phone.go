package phone

import (
    "fmt"
    "log"

    //"github.com/sundowndev/phoneinfoga/v2/lib/number"
    //"github.com/sundowndev/phoneinfoga/v2/lib/remote"
    //"github.com/sundowndev/phoneinfoga/tree/master/lib/remote"
    "gosint/lib/phone/lib/number"
    "gosint/lib/phone/lib/remote"
)

func Phone(phone_num string) {
    n, err := number.NewNumber(phone_num)
    if err != nil {
        log.Fatal(err)
    }

    res, err := remote.NewGoogleSearchScanner().Scan(n)
    if err != nil {
        log.Fatal(err)
    }

    links := res.(remote.GoogleSearchResponse)
    for _, link := range links.Individuals {
        fmt.Println(link.URL) // Google search link to scan
    }
}

/*func main() {
   phone("+33148440709") 
}*/

