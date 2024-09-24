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
    file, err := os.Open(filepath)
    if(err != nil) {
        fmt.Println("Failed to open json file for encoding")
        return
    }

    // Add session to history 
    session_history = append(session_history, session)

    // Convert session history to json data
    writeData, err := json.MarshalIndent(session_history, "", " ")

    // Write json data into history file 
    err = file.Truncate(0)
    if(err != nil) {
        fmt.Println("Failed to truncate json file")
        return
    }

    _, err = file.Write(writeData)
    if(err != nil) {
        fmt.Println("Failed to write data to json file")
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
    fmt.Println(string(byte_value))
    fmt.Println(session_history[0])
    if(len(session_history) == 1) {
        fmt.Println("yes")
    }
    return session_history
}
