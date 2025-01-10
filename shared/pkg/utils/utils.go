package utils

import "log"

// LogError logs an error message.
func LogError(err error, message string) {
    if err != nil {
        log.Printf("%s: %v", message, err)
    }
}