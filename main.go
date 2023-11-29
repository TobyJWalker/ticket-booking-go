package main // Main package

import "fmt" // Import fmt package for printing etc

// Main function
func main() {

	// Initialise globals
	const CONFERENCE_NAME = "Go Conference"
	const CONFERENCE_TICKETS int = 50

	// Initialise variables
	var remainingTickets uint = 50

	// Display value data types
	fmt.Printf("\n\nConference name is of type: %T\n", CONFERENCE_NAME)
	fmt.Printf("Conference tickets is of type: %T\n", CONFERENCE_TICKETS)
	fmt.Printf("Remaining tickets is of type: %T\n\n", remainingTickets)

	// Display welcome message and ticket information
	fmt.Printf("Welcome to %s booking application.\n", CONFERENCE_NAME)
	fmt.Printf("There are %d total spaces available and %d remaining.\n", CONFERENCE_TICKETS, remainingTickets)
	fmt.Println("Get your tickets here to attend our conference!")

	// Initialise booking slice
	var bookings []string // specify a number in the [] to create an array (fixed size)

	// Initialise user input variable
	var firstName string
	var lastName string
	var emailAddress string
	var userTickets uint

	// loop to keep asking for bookings
	for {

		// Get user input
		fmt.Print("\nPlease enter your first name: ")
		fmt.Scanln(&firstName) // & is used to get the memory address of the variable

		fmt.Print("Please enter your last name: ")
		fmt.Scanln(&lastName) 

		fmt.Print("Please enter your email address: ")
		fmt.Scanln(&emailAddress) 

		fmt.Print("Please enter the number of tickets you would like to purchase: ")
		fmt.Scanln(&userTickets)

		// Calculate remaining tickets
		remainingTickets = remainingTickets - userTickets

		// Add name to bookings slice
		bookings = append(bookings, firstName + " " + lastName)


		// Display user input
		fmt.Printf("\nThank you %s, you have booked %d tickets.\n", firstName, userTickets)
		fmt.Printf("A confirmation email has been sent to %s.\n", emailAddress)

		// Display remaining tickets
		fmt.Printf("\nThere are %d tickets remaining.\n", remainingTickets)

		// Display amount of bookings
		fmt.Printf("\nThere are %d bookings.\n", len(bookings))

		// Display all bookings
		fmt.Println("\nBookings:")
		fmt.Printf("%s\n", bookings)
		
	}
}