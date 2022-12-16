package main

import (
	"fmt";
	"reflect"
)


func main(){
	fmt.Println("Insert a valid account : ")
	var account string
	fmt.Scanln(&account)
	var account2 = "dflkdfklfj"
	if reflect.TypeOf(account) == reflect.TypeOf(account2){
		fmt.Printf("%s is a String\n", account)
	}
}