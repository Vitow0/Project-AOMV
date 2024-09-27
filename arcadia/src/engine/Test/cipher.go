package main                       // This program is unoficial and need to be put in test for a new sprite :

import (
	"fmt"
)

func Talk(s string) []int {
	var result []int // create a variable int
	for _, char := range s { // take all the character
		result = append(result, int(char)) // convert in int
	}
	return result // return the variable 
}

func main() {
	message := "This game suck!"
	Cipher := Talk(message) // use the fonction talk to encrypted
	fmt.Println("Message original:", message) // print the original message
	fmt.Println("Message chiffré:", Cipher) // print the crypted message
}