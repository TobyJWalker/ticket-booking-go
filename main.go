package main // Main package

import "fmt" // Import fmt package for printing etc

// Main function
func main() {

	// Initialise globals
	const CONFERENCE_NAME = "Go Conference"
	const CONFERENCE_TICKETS int8 = 50

	// Initialise variables
	var remainingTickets int8 = CONFERENCE_TICKETS

	// Display value data types
	fmt.Printf("\n\nConference name is of type: %T\n", CONFERENCE_NAME)
	fmt.Printf("Conference tickets is of type: %T\n", CONFERENCE_TICKETS)
	fmt.Printf("Remaining tickets is of type: %T\n\n", remainingTickets)

	// Display welcome message and ticket information
	fmt.Printf("Welcome to %s booking application.\n", CONFERENCE_NAME)
	fmt.Printf("There are %d total spaces available and %d remaining.\n", CONFERENCE_TICKETS, remainingTickets)
	fmt.Println("Get your tickets here to attend our conference!")

	// Initialise user input variable
	var userName string = ""
	var userTickets int = 0

	// Get user input
	fmt.Print("\nPlease enter your name: ")
	fmt.Scanln(&userName)

	fmt.Print("Please enter the number of tickets you would like to purchase: ")
	fmt.Scanln(&userTickets)

	// Display user input
	fmt.Printf("Thank you %s, you have purchased %d tickets.\n", userName, userTickets)

}