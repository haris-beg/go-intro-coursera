package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
Write a program which prompts the user to enter integers and
stores the integers in a sorted slice.
The program should be written as a loop.
Before entering the loop, the program should
create an empty integer slice of size (length) 3.
During each pass through the loop, the program prompts the user
to enter an integer to be added to the slice.
The program adds the integer to the slice, sorts the slice,
and prints the contents of the slice in sorted order.
The slice must grow in size to accommodate any number of integers
which the user decides to enter.
The program should only quit (exiting the loop) when the user
enters the character ‘X’ instead of an integer.
*/
func main() {
    const initLen int = 3
    ctr := 0

    // create an empty integer slice of size (length) 3
    slice := make([]int, initLen)

    // setup a loop to accept user input
    for {
        var input string
        fmt.Print("Enter an integer (or X to end): ")
        num, err := fmt.Scan(&input)
        if err!=nil || num != 1{
            fmt.Println("Invalid input!!!")
            continue
        }

        // remove all leading and trailing whiespace from user input, just in case
        input = strings.TrimSpace(input) // O(n)|O(n)

        // if user entered 'X' then end the loop
        if input == "X" {
            break
        }

        i, err := strconv.Atoi(input) // convert user input to integer
        if err!=nil {
            fmt.Println("Invalid input!!!")
            continue
        }

        // adds the integer to the slice
        if ctr < initLen {
            slice[ctr] = i
            ctr++
        } else {
            slice = append(slice, i) // O(1)|O(n)
        }

    }
    sort.Ints(slice) // sorts the slice O(logN)
    fmt.Println(slice) // prints the contents of the slice in sorted order
}