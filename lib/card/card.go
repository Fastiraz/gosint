package card

import (
	"fmt"
	"regexp"
	"strconv"
)

func Card(cardNumber string) {
	// Remove any non-numeric characters from the input
	cardNumber = removeNonNumeric(cardNumber)

	// Check if the card number is valid
	if validateCardNumber(cardNumber) {
		// Identify and print the card type
		cardType := identifyCardType(cardNumber)
		fmt.Printf("Card Type: %s\n", cardType)

		// Extract and print the Issuer Identification Number (IIN/BIN)
		iin := extractIIN(cardNumber)
		fmt.Printf("Issuer Identification Number (IIN/BIN): %s\n", iin)

		// You can add more functions to extract additional information from the card number if needed
		// Remove any non-numeric characters from the input
		cardNumber = removeNonNumeric(cardNumber)
		// extractCreditCardInfo(cardNumber)
	} else {
		fmt.Println("Invalid card number. Please check and try again.")
	}
}

// Function to identify the card type based on the first digit(s) of the card number
func identifyCardType(cardNumber string) string {
	if cardNumber[0] == '4' {
		return "Visa"
	} else if cardNumber[:2] == "34" || cardNumber[:2] == "37" {
		return "American Express"
	} else if cardNumber[0] == '5' {
		return "Mastercard"
	} else if cardNumber[0] == '6' {
		return "Discover"
	} else {
		return "Unknown card type"
	}
}

// Function to extract the issuer identification number (IIN/BIN) from the card number
func extractIIN(cardNumber string) string {
	return cardNumber[:6] // Adjust the length (6 or 8) based on the card type
}

// Function to validate the credit card number using the Luhn Algorithm
func validateCardNumber(cardNumber string) bool {
	sum := 0
	alternate := false

	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(cardNumber[i]))

		if alternate {
			digit *= 2
			if digit > 9 {
				digit = (digit % 10) + 1
			}
		}

		sum += digit
		alternate = !alternate
	}

	return sum%10 == 0
}

// Function to remove any non-numeric characters from the input
func removeNonNumeric(input string) string {
	re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(input, "")
}
