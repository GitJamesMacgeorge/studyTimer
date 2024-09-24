package main

import (
    "fmt"
    "strings"
)

func handleOptionInput(user_input) {
    switch {
        case strings.TrimSpace(user_input) == "1":
            startStudySession()
    }
}

func loadOption(option_number int, show_all bool) string {
    if(show_all) {
        return "1. Start Study Session" + "\n" + "2. View History" + "\n" + "3. Settings" + "\n" +"4. Exit"
    }
     
    switch {
        case option_number == 1:
            return "1. Start Study Session"
        case option_number == 2:
            return "2. View History"
        case option_number == 3:
            return "3. Settings"
        case option_number == 4:
            return "4. Exit"
    }

    return "failed"
}

func menu() {
    fmt.Println(loadOption(0, true))
    for(true) {
        var user_input string 
        fmt.Println("Enter Option: ")
        fmt.Scan(&user_input)
        handleOptionInput(user_input)
    }

}

func welcome() {
    // Show welcome message
    fmt.Println("Welcome")
}

func main() {
    welcome()
    }
