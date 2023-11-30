package main

import (
	"fmt"     // Import fmt package for printing etc
	"strings" // Import strings package for string manipulation
	"time"    // Import time package for time manipulation
)

// Function to validate user input, tell them what is wrong and return a boolean
func isValidInput(firstName string, lastName string, emailAddress string, userTickets int) bool {

	// Name validation, must be at least 2 characters long
	isInvalidName := len(firstName) < 2 || len(lastName) < 2

	if isInvalidName {
		fmt.Printf("\nSorry, your name must be at least 2 characters long.\n")
		fmt.Printf("Please try again.\n")
		return false
	}

	// Email validation, must contain an @ symbol
	isInvalidEmail := !strings.Contains(emailAddress, "@")

	if isInvalidEmail {
		fmt.Printf("\nSorry, your email address must contain an @ symbol.\n")
		fmt.Printf("Please try again.\n")
		return false
	}

	// Ticket validation, must be at least 1 ticket
	isInvalidTicketNumber := userTickets < 1

	if isInvalidTicketNumber {
		fmt.Printf("\nSorry, you must book at least 1 ticket.\n")
		fmt.Printf("Please try again.\n")
		return false
	}

	// Check if there are enough tickets remaining
	if uint(userTickets) > remainingTickets {
		fmt.Printf("\nSorry, there are only %d tickets remaining.\n", remainingTickets)
		fmt.Printf("Please try again.\n")
		return false
	}

	return true
}

// Function to display bookings
func displayBookings(){

	// Get a list of the first names of bookings (for privacy)
	firstNames := []string{}

	for _, booking := range bookings { // _ is used to ignore the index value given with range

		// Get the first name of the booking and add to slice
		firstNames = append(firstNames, booking.first_name)
		
	}

	// Display list of booking first names
	fmt.Println("\nBookings:")
	fmt.Printf("%s\n", firstNames)
}

// Function to simulate sending a ticket
func sendTicket(booking UserData) {

	// Create ticket string
	var ticket string = fmt.Sprintf("%d tickets booked for %s %s.", booking.ticket_count, booking.first_name, booking.last_name)

	// Simulate loading time
	time.Sleep(20 * time.Second) // 2 seconds (number * time unit)

	// Display ticket information
	fmt.Println("\n#############################################")
	fmt.Printf("\nSending ticket to %s.\n\n%s", booking.email_address, ticket)
	fmt.Println("\n\n#############################################")

	// Mark thread as done
	wg.Done()
}