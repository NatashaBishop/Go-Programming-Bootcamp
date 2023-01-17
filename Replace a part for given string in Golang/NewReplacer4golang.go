package main

import "fmt"
import	"strings"

func main() {
	r := strings.NewReplacer("(", "easy,", ")", "tough :-)") // setting the rules: what to replace with what
	fmt.Println(r.Replace("Learning Golang is ( not )"))     //Feed the r function (responcible for the replacement rules) with the string that needs to be amended
}
