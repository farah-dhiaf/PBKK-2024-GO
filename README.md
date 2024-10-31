# PBKK-2024-GO

Nama : Farah Dhia Fadhila </br>
NRP : 5025211030

This repository is for Framework Programming (PBKK) course assignment. Will be updated further for next assignments.

---
## Introduction
Go adalah bahasa pemrograman yang dibangun oleh Google pada 2007 dan menjadi open-source pada 2009. Go dibangun dengan tujuan untuk menyempurnakan bahasa pemrograman yang ada, seperti C, Phyton dan yang lainnya. Go dirancang untuk dapat berjalan pada multiple cores dan menjadi bahasa pemrograman yang sudah built-in mekanisme concurrency.

## Characteristics of Go
- Memiliki syntax yang simpel dan mudah dibaca dari bahasa pemrograman dinamis seperti pada bahasa Python.
- Memiliki efisiensi dan kecepatan dari bahasa pemrograman lower-level dan statis layaknya C.
- Digunakan sebagai bahasa pemrograman server-side atau backend.

## Installation and Setup
1. Install bahasa Go pada https://golang.org/dl/.
2. Setelah selesai proses instalasi, pada Windows, jalankan file `.msi` dab mengikuti wizard instalasi sampai selesai.
3. Lakukan Go modules untuk mengelola dependensi eksternal atau library dengan command di bawah. Command di bawah akan secara otomatis membuat file `go.mod` yang akan menyimpan segala dependensi dari proyek ini.
```
go mod init <nama-modul>
```

## Section 1
Pada bagian ini mencakup introduction dan proses instalasi dari Go seperti yang telah dijelaskan sebelumnya. Selain itu, pada section 1 ini juga membahas terkait:
1. Dasar-dasar dari bahasa Go seperti syntax, variable, data types, pointer, loops dan conditionals.
2. Pada section ini juga membahas terkait logic dari booking tickets, bagaimana user input, dan user validation.

Berikut adalah kode pada section 1:
```
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

```

## Section 2
Pada dasarnya section 2 adalah bagian lanjutan dari section 1, yang mana pada section ini membahas:
1. Encapsulate logic with function untuk mengatur penulisan kode sehingga mudah untuk dibaca dan lebih terlihat ringkas.
2. Go packages untuk mengatur kode ke dalam package-package antar files dan direktori. Tujuannya adalah untuk memudahkan dalam maintaining aplikasi.
3. Scope rules
   - Local scope, variabel yang didefinisikan di dalam fungsi dan hanya bisa diakses di dalam fungsi tersebut.
   - Package scope, variabel yang didefinisikan di luar fungsi dan bisa diakses oleh semua fungsi di dalam package yang sama.
   - Global scope, variabel yang diekspor dari satu package dan bisa diakses oleh package yang lain.
5. Maps, tipe data yang menyimpan key-value.
6. Struct, tipe data yang memungkinkan menyimpan banyak value dengan tipe data yang berbeda.
7. Goroutines, untuk menjalankan concurrency.

Berikut adalah kode pada section 2: </br>
Pada `main.go`
```
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
```
Pada file `helper.go`
```
package helper
import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool){
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email,"@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
	// in Go you can return more than 1 value
}
```
