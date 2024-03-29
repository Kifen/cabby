package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	"github.com/Kifen/cabby/myapp"
)

func main() {
	drop := myapp.Cabby{}
	drop.InitValues()

	cabbyDestinations := drop.GetDestinations()
	fmt.Println("Welcome to cabby.\nWe ply the following destinations: ")
	fmt.Println(cabbyDestinations)

	pickUpPoint, dropOffPoint := start(drop)

	drop.SetDropOffPoint(dropOffPoint)
	drop.SetPickUpPoint(pickUpPoint)
	drop.SetStartTime(time.Now())
	tFare := drop.CalculateFare(pickUpPoint, dropOffPoint)

	tFare = (math.Floor(tFare*100) / 100)
	drop.SetTfare(tFare)

	fmt.Printf("\nIt will cost ₦%.2f to go from %v to %v\n", tFare, pickUpPoint, dropOffPoint)
	time.Sleep(10 * time.Second)
	val, amountToPay := collectFare(drop)

	if val == 1 {
		change := amountToPay - tFare
		fmt.Printf("\nHere's your change of ₦%.2f\n", change)
	}

	fmt.Println("I would appreciate a tip...\n")
	tip, msg := requestTip(tFare)
	fmt.Println(msg)

	drop.SetTip(tip)
	drop.SetEndTime(time.Now())
	drop.TripDetails()
}

func destValid(val bool, spot int) {
	if val == false && spot == 1 {
		fmt.Println("\nInvalid Pickup Destination...")
		fmt.Println("Bye laters...")
		os.Exit(1)
	} else if val == false && spot == 2 {
		fmt.Println("\nInvalid Dropoff Destination...")
		fmt.Println("Bye laters...")
		os.Exit(1)
	}
}

func collectFare(drop myapp.Cabby) (int, float64) {
	i := -1
	var amountToPay float64
	count := 0
	fmt.Println("Enter amount to pay:")

	for i < 0 {
		_, err := fmt.Scan(&amountToPay)
		if err != nil {
			fmt.Println("\nYou entered an invalid value...")
			fmt.Println("Re-enter amount to pay:")
			continue
		}

		i = drop.CheckUserFare(amountToPay)

		if i == -1 {
			fmt.Println("Insufficient funds...")
			fmt.Println("Re-enter amount to pay:")
		}

		count++

		if count%5 == 0 {
			fmt.Println("\nYou'll be reported to the police if you keep on trying to pay less than what you owe\n")
			fmt.Println("Re-enter amount to pay:")
			continue
		}
	}
	return i, (math.Floor(amountToPay*100) / 100)
}

func requestTip(fare float64) (float64, string) {
	var tip float64
	var msg string
	loop := 1
	fmt.Println("Enter your tip:")

	for loop == 1 {
		_, err := fmt.Scan(&tip)
		if err != nil {
			fmt.Println("So u no wan even give me small thing abi...\nAbeg show love:")
			loop = 1
			continue
		}

		if tip <= 0 {
			msg = "Oga u stingy sha ohh...\nBye bye"
			loop = -1
			tip = 0
		} else if tip > 0 && tip <= fare {
			msg = "Thank you!!!"
			loop = -1
		} else {
			msg = "Gracias Mucho!!!!!"
			loop = -1
		}
	}
	return (math.Floor(tip*100) / 100), msg
}

func start(drop myapp.Cabby) (string, string) {
	i := 0
	itr := 1
	var pickUpPoint, dropOffPint string
	for i == 0 {
		pickUpDest := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your pickup destination:")

		pud, _ := pickUpDest.ReadString('\n')
		pickUpPoint = strings.ToLower(strings.Trim(pud, " \r\n"))
		isPudValid := drop.DestinationIsValid(pickUpPoint)
		destValid(isPudValid, 1)

		dropOffDest := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your drop off destination:")

		dod, _ := dropOffDest.ReadString('\n')
		dropOffPint = strings.ToLower(strings.Trim(dod, " \r\n"))
		isDodValid := drop.DestinationIsValid(dropOffPint)
		destValid(isDodValid, 2)
		itr++

		if itr%5 == 0 {
			fmt.Println("\nIt looks like you came here to play.\nHave a good day...\n")
			fmt.Println("Bye laters...")
			os.Exit(1)
		}

		if pickUpPoint == dropOffPint {
			fmt.Println("\nDrop off point and pick up point cannot be same...\n")
			i = 0
			continue
		} else {
			i = -1
		}
	}
	return pickUpPoint, dropOffPint
}
