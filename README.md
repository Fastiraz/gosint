# GOSInt

![GOSInt](./img/GOSInt-removebg-preview.png)

This is a Go script that reads a JSON file containing a list of websites and URLs, prompts the user for a username, replaces the `{}` in the URLs with the username, and makes an HTTP GET request to each URL to check if the username exists.

>**Warning**
>It's not working well, i just made it for learning Golang :)

The script prints the name of each website and the status code and message of the HTTP response for each URL in different colors based on the status code:
- Green: Status code 2xx (successful response)
- Yellow: Status code 4xx (client error)
- Red: Status code 5xx (server error)

## Usage

To use the script, follow these steps:
1. Clone this repository to your local machine
2. Navigate to the root directory of the repository
3. Run the script using the following command:

```bash
go run main.go
```

4. Enter a username when prompted

The script will then make an HTTP GET request to each URL in the JSON file with the `{}` replaced by the username and print the result for each URL.

## Example

Suppose we have the following JSON file named `data.json`:

```json
{
 "7Cups": {
     "errorType": "status_code",
     "url": "https://www.7cups.com/@{}",
     "urlMain": "https://www.7cups.com/",
     "username_claimed": "blue"
 },
 "8tracks": {
     "errorMsg": "This page has vanished",
     "errorType": "message",
     "url": "https://8tracks.com/{}",
     "urlMain": "https://8tracks.com/",
     "username_claimed": "blue"
 },
 "9GAG": {
     "errorType": "status_code",
     "url": "https://www.9gag.com/u/{}",
     "urlMain": "https://www.9gag.com/",
     "username_claimed": "blue"
 }
}
```

And suppose we run the script and enter the username `example_user` when prompted. The output might look like this:

```bash
7Cups: 200 OK
8tracks: 404 Not Found
9GAG: 404 Not Found
Requests completed successfully.
```

In this example, the 7Cups URL returned a successful response with a status code of 200, while the 8tracks and 9GAG URLs both returned a client error with a status code of 404.
