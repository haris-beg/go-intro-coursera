/*
Write a program which prompts the user to first enter a name,
and then enter an address. Your program should create a map
and add the name and address to the map using the keys “name”
and “address”, respectively. Your program should use Marshal()
to create a JSON object from the map, and then your program
should print the JSON object.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
    // prompts the user to enter a name
    fmt.Print("Please enter a name: ")

    consoleScanner := bufio.NewScanner(os.Stdin)
    consoleScanner.Scan()
    name := consoleScanner.Text()

    // prompts the user to enter an address
    fmt.Print("Please enter an address: ")
    consoleScanner.Scan()
    address := consoleScanner.Text()

    // create a map and add the name and address to the map 
    // using the keys “name” and “address”, respectively
    myMap := make(map[string]string)
    myMap["name"] = name
    myMap["address"] = address

    // use Marshal() to create a JSON object from the map
    jsonByteArray, err := json.Marshal(myMap)
    if err != nil {
        fmt.Println(err)
    }

    // print the JSON object
    fmt.Print("You entered: ")
    _, err1 := os.Stdout.Write(jsonByteArray)
    if err1 != nil {
        fmt.Println(err1)
    }
}