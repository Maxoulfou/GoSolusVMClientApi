package command

import "fmt"

func Reboot(array []string) {
	fmt.Printf("Status: %+v\nStatus MSG: %+v\nHostname: %+v\nIP: %+v\n", array[0], array[1], array[2], array[3])
}

func Boot(array []string) {
	fmt.Printf("Status: %+v\nStatus MSG: %+v\nHostname: %+v\nIP: %+v\n", array[0], array[1], array[2], array[3])
}

func Shutdown(array []string) {
	fmt.Printf("Le serveur rompiche\nStatus: %+v\nStatus MSG: %+v\nHostname: %+v\nIP: %+v\n", array[0], array[1], array[2], array[3])
}

func Status(array []string) {
	fmt.Printf("Status: %+v\nStatus MSG: %+v\nVM Stat: %+v\nHostname: %+v\nIP: %+v\n", array[0], array[1], array[2], array[3], array[4])
}

func Info(array []string) {
	fmt.Printf("Status: %+v\nHostname: %+v\nIP: %+v\n", array[0], array[2], array[3])
}
