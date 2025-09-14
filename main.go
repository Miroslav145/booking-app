package main

import "fmt"

func main() {

	// typing
	conferenceName := "Go conference"
	const conferenceTickets int = 50
	remainingTickets := 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %s booking application\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	/*Array: In Go, arrays have fixed sizes
	Example:var bookings = [50]string{"Kylian", "Javier", "Marco"} */
	var bookings [50]string

	//explicit typing
	var firstName string
	var lastName string
	var email string
	var userTickets int

	//  Ask users for their name
	fmt.Print("\nEnter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets -= userTickets

	bookings[0] = firstName + " " + lastName
	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Array type: %T\n", bookings)
	fmt.Printf("Array length: %d\n", len(bookings))

	fmt.Printf("\nThank you %s %s for booking %d tickets. You will receive a confirmation email at %s\n", firstName, lastName, userTickets, email)
	fmt.Printf("%d tickets remaining for %s", remainingTickets, conferenceName)

}
