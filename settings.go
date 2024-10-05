package main

import (
    "fmt"
    "strconv"
    "io/ioutil"
    "strings"
)

func changeDefaultTimer(option string, study_mins *int, break_mins *int) {
    user_reply := promptUser("New Default Timer: ", nil)
    _, err := strconv.Atoi(user_reply)
    if(err != nil) {
        fmt.Println("Invalid Argument for Default Study Time")
        return    
    }

    // Change the timer 
    if(option == "1") {
        // Change Study Mins Timer 
        *study_mins, err = strconv.Atoi(user_reply)
    } else if(option == "2"){
        *break_mins, err = strconv.Atoi(user_reply)
    }
}

func extractSettings(filepath string) map[string] int{
    settings_raw, err := ioutil.ReadFile(filepath)
    if(err != nil) {
        fmt.Println("Reading Settings file failed")
        return nil
    }
    
    studyTime, err := strconv.Atoi(strings.Split(strings.Split(string(settings_raw), "\n")[0], "=")[1])
    breakTime, err := strconv.Atoi(strings.Split(strings.Split(string(settings_raw), "\n")[1], "=")[1])
    fmt.Println(breakTime)
    settings_map := map[string] int {
        "studyTime": studyTime,
        "breakTime": breakTime,
    }

    return settings_map
}

func displaySettings(study_mins *int, break_mins *int) {
    settings_map := extractSettings("settings.txt")
    fmt.Printf("Settings:\n")
    fmt.Printf("Default Study Minutes: %d\n", *study_mins)
    fmt.Printf("Default Break Minutes: %d\n", *break_mins)
    fmt.Println(settings_map)
    // Change Settings
    fmt.Printf("\n1. Change Study Timer\n2. Change Break Timer\n3. Exit\n")
    user_reply := promptUser("Enter Option: ", []string{"1", "2", "3"})
    if(user_reply == "3") {
        return
    }

    changeDefaultTimer(user_reply, study_mins, break_mins)
}


