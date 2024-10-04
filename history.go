package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

type Session struct {
    StudyMinutes int `json:"StudyMinutes"`
    BreakMinutes int `json:"BreakMinutes"`
    Date string `json:"Date"` 
}

func saveSession(filepath string, session_history []Session, session Session) { 
    // Saves a sesson to the json file
    file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
    if(err != nil) {
        fmt.Println("Failed to open json file for encoding")
        return
    }

    defer file.Close()

    // Add session to history 
    session_history = append(session_history, session)

    // Convert session history to json data
    writeData, err := json.MarshalIndent(session_history, "", " ")

    // Write json data into history file 
    err = os.Truncate(filepath, 0) 
    if(err != nil) {
        fmt.Println("Failed to truncate json file")
        fmt.Println(err)
        return
    }

    _, err = file.Write(writeData)
    if(err != nil) {
        fmt.Println("Failed to write data to json file")
        fmt.Println(err)
        return
    }

    fmt.Println("The study session has been saved.")
}

func retrieveSessions(filepath string) []Session {
    // Check if file exists
    file, err := os.Open(filepath)
    if(err != nil) {
        fmt.Println("Failed to open json file for decoding")
        return nil
    }
    defer file.Close()

    // Read sessions stored in json file 
    byte_value, err := ioutil.ReadAll(file)
    if(err != nil) {
        fmt.Println("Failed to read json file")
        return nil
    }

    // Unpack the data into a session struct
    var session_history []Session

    err = json.Unmarshal(byte_value, &session_history)
    if(err != nil) {
        fmt.Println("Failed to Unmarshal json file")
        fmt.Println(err)
        return nil
    }

    if(len(session_history) == 1) {
        fmt.Println("yes")
    }

    return session_history
}

func displayHistory(session_history []Session) {
    for index, value := range session_history {
        fmt.Printf("Session ID: %d\n", index)
        fmt.Printf("Session Date: %s\n", value.Date)
        fmt.Printf("Study Minutes: %d\n", value.StudyMinutes)
        fmt.Printf("Break Minutes: %d\n\n", value.BreakMinutes)
    }
}
