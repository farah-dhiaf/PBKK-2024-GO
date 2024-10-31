package main

import (
	"fmt"
	"strings"
)

func main() {
	// print("Hello, World!")
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint= 50
	bookings := []string{} //array of 50 strings

	// DATA TYPE
	fmt.Printf("conference tickets is %T, remaining tickets is %T, conference name is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We have a total of : %v tickets and %v tickets available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	
	for true{
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		
		// POINTER
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName) //pointer pointing to the memory address of the variable
		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email:")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email,"@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		// CONDITIONAL
		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName + " " + lastName)
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)

			// ARRAY
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				//end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Invalid name. Your name is too short")
			  }
			  if !isValidEmail {
				fmt.Println("Invalid email. Email must contain an '@' sign")
			  }
			  if !isValidTicketNumber {
				fmt.Printf("Invalid number of tickets.")
			  }
			  continue
		}
	}

	//// SCLICE
	// fmt.Printf("The whole slice: %v\n", bookings) //all elements
	// fmt.Printf("The first slice: %v\n", bookings[0]) //first element
	// fmt.Printf("Slice type: %T\n", bookings) //data type
	// fmt.Printf("Slice length: %v\n", len(bookings)) //length
	
	// conferenceTickets = 100 // This will throw an error
	// fmt.Println("Tickets left:", conferenceTickets)

	
}
