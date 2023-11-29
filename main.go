package main // Main package

import "fmt" // Import fmt package for printing etc

// Main function
func main() {

	// Initialise globals
	const CONFERENCE_NAME string = "Go Conference"
	const CONFERENCE_TICKETS int = 50

	// Initialise variables
	var remainingTickets int = CONFERENCE_TICKETS

	// Display welcome message and ticket information
	fmt.Printf("Welcome to %s booking application.\n", CONFERENCE_NAME)
	fmt.Printf("There are %d total spaces available and %d remaining.\n", CONFERENCE_TICKETS, remainingTickets)
	fmt.Println("Get your tickets here to attend our conference!")

}