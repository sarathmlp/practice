package main

import (
    "fmt"
)

func main() {
    var start string
    route := make(map[string]string)
    reverse := make(map[string]string)

    route["Dubai"] = "Russia"
    route["India"] = "Singapore"
    route["Russia"] = "Japan"
    route["Singapore"] = "Dubai"

    for key, val := range route {
        reverse[val] = key
    }
    for key, _ := range route {
        if _, found := reverse[key]; !found {
            start = key
        }
    }
    for {
        var val string
        var found bool

        if val, found = route[start]; found {
            fmt.Print(start, "->", val, ", ")
        } else {
            break
        }
        start = val
    }
}
