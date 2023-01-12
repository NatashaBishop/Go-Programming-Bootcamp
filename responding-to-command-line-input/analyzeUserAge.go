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
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please, input your age:")
	age, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	age = strings.TrimSpace(age)
	iAge, err := strconv.Atoi(age)
	if err != nil {
		log.Fatal(err)
	}
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
