package main

import (
	"fmt"

	"github.com/Vikesh24/cryptit/encrypt"

	"github.com/Vikesh24/cryptit/decrypt"
)

func main(){
	var userInput string
	fmt.Print("Please enter the string: ")
	fmt.Scanf("%s", &userInput)
	encryptedStr := encrypt.Nimbus(userInput)
	decryptedStr := decrypt.Nimbus(encryptedStr)
	fmt.Println(encryptedStr, decryptedStr)
}