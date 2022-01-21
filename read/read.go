/*
Write a program which reads information from a file
and represents it in a slice of structs.
Assume that there is a text file which contains a series of names.
Each line of the text file has a first name and a last name,
in that order, separated by a single space on the line.

Your program will define a name struct which has two fields,
fname for the first name, and lname for the last name.
Each field will be a string of size 20 (characters).
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// define a name struct which has two fields
// fname for the first name, and lname for the last name
// Each field will be a string of size 20 (characters).
type Name struct {
    fname string
    lname string
}

func main() {
    // prompt the user for the name of the text file
    fmt.Print("Enter name of text file to read: ")
    var fileName string
    _, err1 := fmt.Scan(&fileName)
    abortIfError(err1)

    // open the file for read-only
    file, err2 := os.Open(fileName)
    abortIfError(err2)

    scanner := bufio.NewScanner(file)

    namesSlice := []Name{}

    // successively read each line of the text file
    for scanner.Scan() {
        split := strings.Split(scanner.Text(), " ")

        // create a struct which contains the first and last names found in the file
        var name = Name{fname: split[0], lname: split[1]}

        // Each struct created will be added to a slice
        namesSlice = append(namesSlice, name)        
    }


    // After reading all lines from the file
    file.Close()

    // iterate through your slice of structs
    for _, v := range namesSlice {
        // print the first and last names found in each struct
        fmt.Println(v.fname, v.lname)
    }
}

func abortIfError(err error)  {
    if err != nil {
        log.Fatal(err)
    }
}