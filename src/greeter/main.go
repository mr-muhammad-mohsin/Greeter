package main

import (
    "fmt"
    "greetings"
)

func main() {
    name:="Ehtasham Zahoor"
    status:="Dr." 
    message:=greetings.Greet(name, status)
    fmt.Println(message)
}