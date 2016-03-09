package greetings

var Status string
var Name string

/*
    Greet Function takes input the name and status of user and return an greet from a new virtual Assistant
*/
func Greet (name string, status string) string {
    Name = name
    Status = status
    return "Hi! " + Status + " " + Name + ", I am your new Assistant"
}
