package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

// var conferenceName = "Go Conference"
var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets = 50

// var bookings = [50] string {}
// var bookings []string
// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	// // var conferenceName = "Go Conference"
	// conferenceName := "Go Conference"
	// const conferenceTickets = 50
	// var remainingTickets = 50
	// // var bookings = [50] string {}
	// var bookings []string

	greetUsers()

	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T \n", conferenceName, conferenceTickets, remainingTickets)

	// fmt.Println("Welcome to ", conferenceName, " bookung application")
	// fmt.Printf("Welcome to %v bookung application\n", conferenceName)
	// fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	// fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	// fmt.Println("Get your tickets here to attend")

	// for remainingTickets > 0 && len(bookings) < 50 {
	// for {
	// var firstName string
	// var lastName string
	// var email string
	// var userTickets uint

	// // ask user for their name
	// fmt.Print("Enter your first name : ")
	// fmt.Scan(&firstName) //& - pointer
	// fmt.Print("Enter your last name : ")
	// fmt.Scan(&lastName)
	// fmt.Print("Enter your email : ")
	// fmt.Scan(&email)
	// fmt.Print("Enter numbetr of tickets : ")
	// fmt.Scan(&userTickets)

	// isValidName := len(firstName) >= 2 && len(lastName) >= 2
	// isValidEmail := strings.Contains(email, "@")
	// isValidTicketNumber := userTickets > 0 && userTickets <= uint(remainingTickets)

	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, uint(remainingTickets))

	// if userTickets <= uint(remainingTickets) {
	if isValidEmail && isValidName && isValidTicketNumber {
		// remainingTickets = remainingTickets - int(userTickets)
		// // bookings[0] = firstName + " " + lastName
		// bookings = append(bookings, firstName+" "+lastName)

		// fmt.Printf("The whole slice : %v \n", bookings)
		// fmt.Printf("The first value: %v \n", bookings[0])
		// fmt.Printf("The slice type : %T \n", bookings)
		// fmt.Printf("The slice length : %v \n", len(bookings))

		// bookings[0]="damindu"
		// userName = "Damindu"
		// userTickets = 2
		// fmt.Printf("User %v booked %v tickets.\n", firstName, userTickets)
		// fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		// firstNames := []string{}
		// for _, booking := range bookings {
		// 	var names = strings.Fields(booking)
		// 	firstNames = append(firstNames, names[0])
		// }

		// fmt.Printf("These are all our bookings: %v\n", bookings)
		// fmt.Printf("The first names of bookings are: %v\n", firstNames)

		bookTicket(userTickets, firstName, lastName, email)
		wg.Add(1)
		go sendTickets(userTickets, firstName, lastName, email) //go use for concurrebt thread
		// getFirstNames(bookings)
		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		noTicketsRemaining := remainingTickets == 0
		if noTicketsRemaining {
			//end programe
			fmt.Println("Our conference is booked out. Come back next year.")
			// break
		}

	} else {
		// fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)

		if !isValidName {
			fmt.Println("first name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("email address you entered is not contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("numbers of tickets you entered is invalid")
		}
		// fmt.Println("Your input data is invalid, try again")
		// continue
	}
	wg.Wait()

}

// city := "London"

// switch city{
// 	case "New York":
// 	case "Singapore":
// 	case "London" , "Berlin":
// 	case "Mexico City" , "Hong kong":
// 	default:
// 		fmt.Println("No valid city selected")

// }

// }

func greetUsers() {
	fmt.Printf("Welcome to %v booking appliication \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		// firstNames = append(firstNames, names[0]) //array
		// firstNames = append(firstNames, booking["firstName"]) //map
		firstNames = append(firstNames, booking.firstName) //struct
	}
	// fmt.Printf("The first names of bookings are: %v\n", firstNames)
	return firstNames
}

// func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
// 	isValidName := len(firstName) >= 2 && len(lastName) >= 2
// 	isValidEmail := strings.Contains(email, "@")
// 	isValidTicketNumber := userTickets > 0 && userTickets <= uint(remainingTickets)

// 	return isValidName, isValidEmail, isValidTicketNumber
// }

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for their name
	fmt.Print("Enter your first name : ")
	fmt.Scan(&firstName) //& - pointer
	fmt.Print("Enter your last name : ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email : ")
	fmt.Scan(&email)
	fmt.Print("Enter numbetr of tickets : ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - int(userTickets)
	// bookings[0] = firstName + " " + lastName

	//create a map for a user
	// var userData = make(map[string]string)

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	// bookings = append(bookings, firstName+" "+lastName)
	bookings = append(bookings, userData)
	fmt.Printf("List of booking is %v \n", bookings)

	fmt.Printf("User %v booked %v tickets.\n", firstName, userTickets)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Printf("Sending ticket : %v to email address %v", ticket, email)
	wg.Done()
}
