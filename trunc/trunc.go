package main
 
import "fmt"
 
/*
Write a program which prompts the user to enter a floating point number and prints the integer,
which is a truncated version of the floating point number that was entered.
Truncation is the process of removing the digits to the right of the decimal place.
 
Submit your source code for the program, “trunc.go”.
*/
 
func main ()  {
    var x float32
 
    fmt.Printf("Enter a floating point number: ")
 
    num, err := fmt.Scan(&x)
 
    if num != 1 {
        fmt.Println("Invalid input")
    }
 
    i := int(x)
 
    if err != nil {
        fmt.Println("Error occured")
    } else {
        fmt.Println(i)
    }
}