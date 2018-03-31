package main

/*
	Languages:
		de
		en-us
		en
		no-nb
		de-ch
		en-au
		en-gb
		nl
		zh-CN
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/manveru/faker"
)

//prompts []string = {"Name","Email","First line of address"}

//Hold individual's details.
type personalDetails struct {
	name          string
	email         string
	streetAddress string
	townCity      string
	postalZipCode string
	county        string
	country       string
	telephone     string
	mobile        string
}

//Create a slice of struct.
var people []personalDetails

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Enter the number of dummy records ('q' or 'Q' to quit): ")
		scanner.Scan()

		if scanner.Text() == "q" || scanner.Text() == "Q" {
			break
		}

		strconv.Atoi(scanner.Text())

		if _, err := strconv.Atoi(scanner.Text()); err == nil {

			numberOfRecords, _ := strconv.Atoi(scanner.Text())

			fmt.Println()

			for i := 1; i <= numberOfRecords; i++ {
				fake, err := faker.New("en")

				if err != nil {
					panic(err)
				}

				fmt.Printf("Record %d\n", i)
				fmt.Println("---------")
				fmt.Printf("Name:\t\t%s\n", fake.NamePrefix()+" "+fake.Name())
				fmt.Printf("Email:\t\t%s\n", fake.Email())
				fmt.Printf("Street:\t\t%s\n", fake.StreetAddress())
				fmt.Printf("City:\t\t%s\n", fake.City())
				fmt.Printf("Postcode:\t%s\n", fake.PostCode())
				fmt.Printf("State:\t\t%s\n", fake.State())
				fmt.Printf("Country:\t%s\n", fake.Country())
				fmt.Println()
			}

			break

		} else {
			fmt.Println("Entered number is not numeric. Please retry.")
		}
	}

}
