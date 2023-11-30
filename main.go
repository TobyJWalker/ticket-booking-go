package main // Main package

import (
	"fmt"     // Import fmt package for printing etc
	"strings" // Import strings package for string manipulation
)

// Main function
func main() {

	// Initialise globals
	const CONFERENCE_NAME = "Go Conference"
	const CONFERENCE_TICKETS uint = 50

	// Initialise variables
	var remainingTickets uint = 50

	// Display value data types
	// fmt.Printf("\n\nConference name is of type: %T\n", CONFERENCE_NAME)
	// fmt.Printf("Conference tickets is of type: %T\n", CONFERENCE_TICKETS)
	// fmt.Printf("Remaining tickets is of type: %T\n\n", remainingTickets)

	// Display welcome message and ticket information
	greetUsers(CONFERENCE_NAME, remainingTickets, CONFERENCE_TICKETS)

	// Initialise booking slice
	var bookings []string // specify a number in the [] to create an array (fixed size)

	// Initialise user input variable
	var firstName string
	var lastName string
	var emailAddress string
	var userTickets uint

	// loop to keep asking for bookings whilst tickets are available
	for remainingTickets > 0{

		// Get user input
		fmt.Print("\nPlease enter your first name: ")
		fmt.Scanln(&firstName) // & is used to get the memory address of the variable

		fmt.Print("Please enter your last name: ")
		fmt.Scanln(&lastName) 

		fmt.Print("Please enter your email address: ")
		fmt.Scanln(&emailAddress) 

		fmt.Print("Please enter the number of tickets you would like to purchase: ")
		fmt.Scanln(&userTickets)

		// Name validation, must be at least 2 characters long
		isInvalidName := len(firstName) < 2 || len(lastName) < 2

		if isInvalidName {
			fmt.Printf("\nSorry, your name must be at least 2 characters long.\n")
			fmt.Printf("Please try again.\n")
			continue // continue to the next iteration of the loop
		}

		// Email validation, must contain an @ symbol
		isInvalidEmail := !strings.Contains(emailAddress, "@")

		if isInvalidEmail {
			fmt.Printf("\nSorry, your email address must contain an @ symbol.\n")
			fmt.Printf("Please try again.\n")
			continue // continue to the next iteration of the loop
		}

		// Ticket validation, must be at least 1 ticket
		isInvalidTicketNumber := userTickets < 1

		if isInvalidTicketNumber {
			fmt.Printf("\nSorry, you must book at least 1 ticket.\n")
			fmt.Printf("Please try again.\n")
			continue // continue to the next iteration of the loop
		}

		// Check if there are enough tickets remaining
		if userTickets > remainingTickets {
			fmt.Printf("\nSorry, there are only %d tickets remaining.\n", remainingTickets)
			fmt.Printf("Please try again.\n")
			continue // continue to the next iteration of the loop
		}

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

		// Get a list of the first names of bookings (for privacy)
		firstNames := []string{}

		for _, booking := range bookings { // _ is used to ignore the index value given with range

			// Get the first name of the booking and add to slice
			var splitNames = strings.Split(booking, " ")
			firstNames = append(firstNames, splitNames[0])
			
		}

		// Display list of booking first names
		fmt.Println("\nBookings:")
		fmt.Printf("%s\n", firstNames)
		
	}
}

// Function to greet users
func greetUsers(conferenceName string, remainingTickets uint, totalTickets uint) {
	fmt.Printf("Welcome to %s booking application.\n", conferenceName)
	fmt.Printf("There are %d total spaces available and %d remaining.\n", totalTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend our conference!")
}