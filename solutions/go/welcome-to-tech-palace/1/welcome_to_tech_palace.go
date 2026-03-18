package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
	panic("Please implement the WelcomeMessage() function")
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    msg := strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg  + "\n" + strings.Repeat("*", numStarsPerLine)
    return(msg) 
	panic("Please implement the AddBorder() function")
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    noStars := strings.ReplaceAll(oldMsg, "*", "")
    cleaned := strings.TrimSpace(noStars)

    return cleaned                         
	panic("Please implement the CleanupMessage() function")
}
