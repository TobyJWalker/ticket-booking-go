package main // Main package

import (
	"fmt" // Import fmt package for printing etc
	"time"
)

// Package level constants
const CONFERENCE_NAME = "Go Conference"
const CONFERENCE_TICKETS uint = 50

// Package level variables (not good practice but for the sake of this example)
var remainingTickets uint = 50
var bookings = make([]UserData, 0) // specify a number in the first [] and remove 0 to create an array (fixed size)

// Create a struct to hold user data
type UserData struct{
	first_name string
	last_name string
	email_address string
	ticket_count int
}

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
		if !isValidInput(firstName, lastName, emailAddress, userTickets) {
			continue // continue to the next iteration of the loop
		}

		// Calculate remaining tickets
		calculateremainingTickets(userTickets)

		// Process booking
		processBooking(firstName, lastName, emailAddress, userTickets)

		// Display remaining tickets
		fmt.Printf("\nThere are %d tickets remaining.\n", remainingTickets)

		// Display bookings
		displayBookings()

		
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

	// Create a map for user data
	var userData = UserData {
		first_name: firstName,
		last_name: lastName,
		email_address: emailAddress,
		ticket_count: userTickets,
	}

	// Add name to bookings slice
	bookings = append(bookings, userData)

	// Display user input
	fmt.Printf("\nThank you %s, you have booked %d tickets.\n", firstName, userTickets)
	fmt.Printf("A confirmation email has been sent to %s.\n", emailAddress)

	// Send ticket
	sendTicket(userData)
}

// Function to simulate sending a ticket
func sendTicket(booking UserData) {

	// Create ticket string
	var ticket string = fmt.Sprintf("%d tickets booked for %s %s.", booking.ticket_count, booking.first_name, booking.last_name)

	// Simulate loading time
	time.Sleep(2 * time.Second) // 2 seconds (number * time unit)

	// Display ticket information
	fmt.Println("\n#############################################")
	fmt.Printf("\nSending ticket to %s.\n\n%s", booking.email_address, ticket)
	fmt.Println("\n#############################################")
}