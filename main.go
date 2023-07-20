package main

import (
	"fmt"
	"time"
	"sync"// provides basic synchronization functional
)

	var conferenceName = "Go Conference"
	const comferenceTickets = 50
	var remainigTickets uint = 50
	var bookings = make([]UserData,0)//a list of user data structs

	/* IF it's MAP:
  var bookings = make([]map[string]string,0)

	creating an empty list of maps, not the emty map. Thats why we have squared braces here. zero is for initializing a size of the list which will grow anyway, so can put any number because the slices are dynamic
  /*

	/*STRUCT:
	mixed data type
	*/
	
	//type is creating a type based on struct of...
	type UserData struct {
		firstName string
		lastName string
		email string
		numberOfTickets uint
	}

	var wg = sync.WaitGroup {

	}

	func main() {
	
			greetUser()

			userName, lastName, email, userTickets := getUserInput()
			isValidEmail, isValidName, isValidTicketNumber := validateUserInput(userName, lastName, email, userTickets)

			if isValidName && isValidEmail && isValidTicketNumber {

				bookTicket(userTickets, userName, lastName, email)

				wg.Add(1)// Sets the number of goroutines(threads) to wait for
				go sendTicket(userTickets, userName, lastName, email)// go starts a new goroutine/thread; 
				firstNames:= getFirstNames()
				fmt.Printf("The first names of bookings are: %v\n", firstNames)

					if remainigTickets == 0 {
						fmt.Println("Our conference is booked out. Come back next year")
						// break
					}

			} else {
				if !isValidTicketNumber {
						fmt.Println("Your tiket amount is invalid. Please try again")
				} 
				if !isValidEmail {
					fmt.Println("Your email must contain @ symbol.Please try again ")
				} 
				if !isValidName {
					fmt.Println("Your first or last name is too short")
				}
		  }
			wg.Wait()// Blocks until the waitGroup counter id 0
		}

	func greetUser() {
		fmt.Printf("Welcome to %v booking aplication\n", conferenceName)
		fmt.Printf("We have in total %v tickets and  %v are availible\n", comferenceTickets, remainigTickets)
	  fmt.Println("Get your tickets here to attend")
	}

	func getFirstNames() []string {
		firstNames := []string{}
		for _, booking := range bookings {
			firstNames = append(firstNames, booking.firstName)
		}
		/* with map: 
		firstNames = append(firstNames, booking["userNames"])
		 **without map**
		 var names = strings.Fields(booking)//fields simply slits element in array by spaces. this func comes from string package thats why we put string. before. in result we will get a slice of stings seperated by space
		*/
		return firstNames
	}

	func getUserInput() (string, string, string, uint) {
		var userName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name:")
		fmt.Scan(&userName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email:")
		fmt.Scan(&email)

		fmt.Println("Enter your number of tikets u need:")
		fmt.Scan(&userTickets)

		return userName, lastName, email, userTickets
	}


	func bookTicket(userTickets uint, userName string, lastName string, email string) {	
		remainigTickets = remainigTickets - userTickets// have to have same data type or need to convert

		//*****************STRUUUUUCTTT**********	
		var userData = UserData {
			firstName: userName,
			lastName: lastName,
			email: email,
			numberOfTickets: userTickets,
		}	
		
		/************************MAAAAAPPPP*************
		create a map for a user:

		map can only have the same data type as keys and the same data type for values 
		var userData = make(map[string]string)

		userData["userName"] = userName
		userData["lastName"] = lastName
		userData["email"] = email
		userData["numberOfTicets"] = strconv.FormatUint(uint64(userTickets),10)//2:46:08 nana| convert a uint value to a string as a decimal number
    */
		bookings = append(bookings, userData)
		fmt.Printf("List of bookings is %v\n", bookings)

		fmt.Printf("Thank you %v %v for booking %v tickets. you will recieve a confirmation on your email %v\n", userName, lastName, userTickets, email)
		fmt.Printf("%v are remaining for booking at %v\n",remainigTickets, conferenceName)
}

func sendTicket(userTickets uint, userName string, lastName string, email string) {
	time.Sleep(10 * time.Second)//stops the execution for the defined time
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, userName, lastName)
	fmt.Println("*****************")
	fmt.Printf("sending ticket:\n %v\n to email adress\n %v\n", ticket, email)
	fmt.Println("****************")
	wg.Done()//Decrements the wait Group counter by 1 So this is called by the goroutine to indicate that it's finished
}
