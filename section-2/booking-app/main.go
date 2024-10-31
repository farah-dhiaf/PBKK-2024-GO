package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50
var conferenceName = "Go Conference"
var remainingTickets uint= 50
// var bookings = make([]map[string]string, 0) //array of 50 strings
var bookings = make([]userData, 0)

// STRUCT allows data with different data types
type userData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}
func main() {

	greetUsers()
	
	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	// CONDITIONAL
	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email) // make it concurrent

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			//end program
			fmt.Println("Our conference is booked out. Come back next year.")
			// break
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
	}
	wg.Wait()
}

//// SCLICE
// fmt.Printf("The whole slice: %v\n", bookings) //all elements
// fmt.Printf("The first slice: %v\n", bookings[0]) //first element
// fmt.Printf("Slice type: %T\n", bookings) //data type
// fmt.Printf("Slice length: %v\n", len(bookings)) //length

// conferenceTickets = 100 // This will throw an error
// fmt.Println("Tickets left:", conferenceTickets)

	

func greetUsers() {
	fmt.Printf("Welcome to %v booking app!\n", conferenceName)
	// fmt.Printf("conference tickets is %T, remaining tickets is %T, conference name is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("We have a total of : %v tickets and %v tickets available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string { //return type is a slice of strings
	// ARRAY
	firstNames := []string{}
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}


func getUserInput()(string, string, string, uint){
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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string){
	remainingTickets = remainingTickets - userTickets

	// MAP
	// var userData = make(map[string]string) //key and value are same data
	var userData = userData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10) //convert uint to string

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string){
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets %v %v\n", userTickets, firstName, lastName) //string interpolation, save to variable and print
	fmt.Println("#################")
	fmt.Printf("Sending ticket to %v to email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}