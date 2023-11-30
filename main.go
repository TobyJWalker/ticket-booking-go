package main // Main package

import (
	"fmt"                      // Import fmt package for printing etc
	"ticket-booking-go/shared" // Import our shared package
)

// Package level constants
const CONFERENCE_NAME = "Go Conference"
const CONFERENCE_TICKETS uint = 50

// Package level variables (not good practice but for the sake of this example)
var remainingTickets uint = 50
var bookings []string // specify a number in the [] to create an array (fixed size)

// Main function
func main() {

	// Display welcome message and ticket information
	greetUsers(CONFERENCE_NAME, remainingTickets, CONFERENCE_TICKETS)

	// Initialise user input variable
	var firstName string
	var lastName string
	var emailAddress string
	var userTickets int

	// loop to keep asking for bookings whilst tickets are available
	for remainingTickets > 0 {

		// Get user input
		firstName, lastName, emailAddress, userTickets = getUserDetails()

		// Validate user input
		if !shared.IsValidInput(firstName, lastName, emailAddress, userTickets, remainingTickets) {
			continue // continue to the next iteration of the loop
		}

		// Calculate remaining tickets
		calculateremainingTickets(userTickets)

		// Process booking
		processBooking(firstName, lastName, emailAddress, userTickets)

		// Display remaining tickets
		fmt.Printf("\nThere are %d tickets remaining.\n", remainingTickets)

		// Display bookings
		shared.Displaybookings(bookings)
		
	}
}

// Function to greet users
func greetUsers(conferenceName string, remainingTickets uint, totalTickets uint) {
	fmt.Printf("Welcome to %s booking application.\n", conferenceName)
	fmt.Printf("There are %d total spaces available and %d remaining.\n", totalTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend our conference!")
}

// Fnction to get user input
func getUserDetails() (string, string, string, int) {

	// Initialise user input variable
	var firstName string
	var lastName string
	var emailAddress string
	var userTickets int

	// Get user input
	fmt.Print("\nPlease enter your first name: ")
	fmt.Scanln(&firstName) // & is used to get the memory address of the variable

	fmt.Print("Please enter your last name: ")
	fmt.Scanln(&lastName) 

	fmt.Print("Please enter your email address: ")
	fmt.Scanln(&emailAddress) 

	fmt.Print("Please enter the number of tickets you would like to purchase: ")
	fmt.Scanln(&userTickets)

	// Return all the data
	return firstName, lastName, emailAddress, userTickets
}

// Function to get remaining tickets
func calculateremainingTickets(userTickets int) {
	remainingTickets = remainingTickets - uint(userTickets)
}

// Function to process booking
func processBooking(firstName string, lastName string, emailAddress string, userTickets int) {

	// Add name to bookings slice
	bookings = append(bookings, firstName + " " + lastName)

	// Display user input
	fmt.Printf("\nThank you %s, you have booked %d tickets.\n", firstName, userTickets)
	fmt.Printf("A confirmation email has been sent to %s.\n", emailAddress)
}