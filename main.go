package main

import "fmt"

func main() {

	// typing
	conferenceName := "Go conference"
	const conferenceTickets int = 50
	remainingTickets := 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %s booking application\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d are stil available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	//explicit typing
	var firstName string
	var lastName string
	var email string
	var userTickets int

	//  Ask users for their name
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets -= userTickets

	fmt.Printf("Thank you %s %s for booking %d tickets. You will receive a confirmation email at %s\n", firstName, lastName, userTickets, email)
	fmt.Printf("%d tickets remaining for %s", remainingTickets, conferenceName)

	//57:15

}
