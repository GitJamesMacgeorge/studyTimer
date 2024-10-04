package main

import (
    "fmt"
    "strings"
    "os"
)

func promptUser(request string, expected_answer []string) string {
   
    for true {
        fmt.Printf("\n%s", request)
        var user_reply string
        fmt.Scan(&user_reply)
        result_con := false

        // Print Request 
       
        for i := 0; i < len(expected_answer); i++ {
            if(strings.ToLower(strings.TrimSpace(user_reply)) == expected_answer[i] && !result_con) {
                result_con = true
                break
            }
        }

        if(result_con) {
            return strings.ToLower(strings.TrimSpace(user_reply))
        }

        return "failed"

    }
    return "failed"
}


func handleOptionInput(session_history []Session, user_input string, study_length_mins int, break_length_mins int) {
    switch {
        case strings.TrimSpace(user_input) == "1":
            session := startStudySession(study_length_mins, break_length_mins)
            saveSession(HISTORY_PATH, session_history, session)
        case strings.TrimSpace(user_input) == "2":
            displayHistory(session_history)
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

func welcome() {
    // Show welcome message
    fmt.Println("Welcome")
}

func main() {
    // Retrieve History
    session_history := retrieveSessions(HISTORY_PATH)
    welcome()
    result := loadOption(0, true)
    if(result == "failed") {
        fmt.Println("LoadOption failed")
        os.Exit(3)
    }
    fmt.Println(result)

    condition := true
    for condition {
        user_reply := promptUser("Enter Option: ", []string{"1", "2", "3", "4"})
        if(user_reply == "failed") {
            fmt.Println("promptUser failed")
            os.Exit(3)
        }
        // Exit app
        if(strings.TrimSpace(user_reply) == "4") {
            os.Exit(3)
        }

        handleOptionInput(session_history, user_reply, 1, 1)
    }
}
