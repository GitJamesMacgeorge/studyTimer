package main

import (
    "fmt"
    "time"
    "strings"
)

func commenceStudyTimer(length_mins int) bool {
    fmt.Println("The timer has started")
    // Set Duration & Create Timer
    duration := time.Duration(length_mins) * time.Minute
    timer := time.NewTimer(duration)
    
    // Waits for the timer to expire
    <-timer.C
    fmt.Println("The timer has finished.")
    return true
}


func showDetails(study_length_mins int, break_length_mins int) bool {
    // Displays the Details of the study session before commencement
    var user_reply string
    fmt.Printf("Study Time: %d\n", study_length_mins)
    fmt.Printf("Break Time: %d\n", break_length_mins)
    
    fmt.Printf("Commence Study? (Y/N): ")
    for true {
        fmt.Scan(&user_reply)
        switch {
            case strings.ToLower(user_reply) == "y":
                return true
            case strings.ToLower(user_reply) == "n":
                return false
            default:
                fmt.Printf("Invalid Answer")
        }
    }
    
    return false
}

func startStudySession(study_length_mins int, break_length_mins int) Session {
    fmt.Println("Starting study session...")
    showDetails(study_length_mins, break_length_mins)

    // Create Session 
    session := Session {
        StudyMinutes: 0,
        BreakMinutes: 0,
        Date: time.Now().Format("2006-01-02"),
    }

    for true {
        studyTimerResult := commenceStudyTimer(study_length_mins)
        session.StudyMinutes += study_length_mins // @@@ Need better

        // Prompt user if they want to begin break timer 
        user_result := promptUser("Do you want to begin the your break? (Y/N): ", []string{"y", "n", "q"})
        if(studyTimerResult && user_result == "n") {
            break
        }
     
        // Begin Break timer & prompt user further 
        breakTimerResult := commenceStudyTimer(break_length_mins)
        user_result = promptUser("Do you want to being commence studying? (Y/N): ", []string{"y", "n", "q"})
        session.BreakMinutes += break_length_mins
        if(breakTimerResult && user_result == "n") {
            break
        }
    }
    
    // Save Study Session
    return session
}
