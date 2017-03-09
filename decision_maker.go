package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "time"
)

func printDashes() {
    fmt.Println("-----------------------------------")
}

func getUserOptions() []string {
    // Declare a slice to hold options user provides
    var options []string
    i := 0
    // Use a loop to prompt and read user input and add to a list
    // of options, if user types 'done' then break loop
    for {
        if i < 1 {
            fmt.Print("Provide an option to choose from: ")
        } else {
            fmt.Print("Provide another option to choose from (type done if complete): ")
        }
        // Create a Scanner to read from user input (os.Stdin)
        scanner := bufio.NewScanner(os.Stdin)
        // Scan advances to token to get usre input then use
        // Text to get value and read it into input
        scanner.Scan()
        input := scanner.Text()
        // If user types exit then exit program and do not make
        // a selection by clearing out any inputted data if it
        // exists
        if input == "exit" {
            fmt.Println("Successfully exited Decision Maker.")
            options = options[:0]
            break
        // If user types done then break out of loop
        } else if input == "done" {
            break
        // Make sure input is being provided
        } else if len(input) == 0 {
            fmt.Println("No option provided, type exit to quit.")
        } else {
            options = append(options, input)
            i++
        }
    }
    return options
}

func makeDecision(options []string) string {
    // Give it a random seed so we can get a different
    // result each run
    rand.Seed(time.Now().UnixNano())
    // Get a choice by passing the len of options slice to
    // the rand.Intn function then use the random value it
    // generates to select an option
    choice := options[rand.Intn(len(options))]
    return choice
}

func main() {
    fmt.Println("Welcome to Decision Maker!")
    fmt.Println("Can't make a decision? Let us help!")
    fmt.Println("Type exit at any time to quit.")
    printDashes()
    // Gather the user options
    options := getUserOptions()
    // If the user provided options then pick one
    if len(options) > 0 {
        choice := makeDecision(options)
        printDashes()
        fmt.Printf("Decision maker has chosen *%s* for you!\n", choice)
        fmt.Println("Thank you for using decision maker.")
    }
}
