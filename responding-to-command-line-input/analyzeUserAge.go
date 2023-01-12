package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Setup buffered reader that gets text from the keyboard
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please, input your age:") //Ask user to input their age
	age, err := reader.ReadString('\n') //Read the input till the new line appears
	if err != nil { //handling error
		log.Fatal(err)
	}
	age = strings.TrimSpace(age) //Get read of white spaces
	iAge, err := strconv.Atoi(age) //Convert ACII input to interger
	if err != nil { //handling error
		log.Fatal(err)
	}
	//Analyze user input and respond according to given logic
	if iAge < 5 {
		fmt.Println("You are too young for school")
	} else if iAge == 5 {
		fmt.Println("It's time to go to kindergarden")
	} else if (iAge > 5) && (iAge <= 17) {
		grade := iAge - 5
		fmt.Printf("You should go to grade %d class", grade) // note using Printf rather than Pringln (formated text printing where we include variable "grade" within the text string)

	} else {
		fmt.Println("It's time to go to college")
	}
}
