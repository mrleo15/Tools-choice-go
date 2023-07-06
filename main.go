package main

import (
	"fmt"
	"os"

	gochoice "github.com/TwiN/go-choice"
	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

func main() {
	// LOGGER
	// logger := log.NewWithOptions(os.Stderr, log.Options{
	// 	ReportCaller:    true,
	// 	ReportTimestamp: true,
	// 	TimeFormat:      time.Kitchen,
	// 	Prefix:          "Baking 🍪 ",
	// })

	// READ ENVIRONMENT VARIABLES
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var pass string
	choice, index, err := gochoice.Pick(
		"What do you want to do?\nPick:",
		[]string{
			"Connect to the production environment",
			"Connect to the test environment",
			"Update",
		},
		gochoice.OptionBackgroundColor(gochoice.Black),
		gochoice.OptionTextColor(gochoice.White),
		gochoice.OptionSelectedTextColor(gochoice.Red),
		// gochoice.OptionSelectedTextBold(),
	)

	if index == 0 {
		fmt.Println("You have selected: 'Connect to the production environment', which is the index 0")
		fmt.Println("Input production password: ")
	} else if index == 1 {
		fmt.Println("You have selected: 'Connect to the test environment', which is the index 1")
		fmt.Println("Input test password: ")
		// input passowrd from cli
		fmt.Scanln(&pass)

		if pass == os.Getenv("TEST_PASSWORD") {
			fmt.Println("Wrong password!")
			return
		}
		for temp := 1; temp <= 150; temp++ {
			log.Info("Read Data tmp", "Prosess", temp)
		}

		fmt.Println("You have inputted: ", pass)
	} else if index == 2 {
		fmt.Println("You have selected: 'Update', which is the index 2")
		fmt.Println("Input Update password: ")
	}
	if err != nil {
		fmt.Println("You didn't select anything!")
	} else {
		fmt.Println("You have selected: ", choice)
	}
}
