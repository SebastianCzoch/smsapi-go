package main

import (
	"fmt"

	"github.com/SebastianCzoch/smsapi-go"
)

func main() {
	client := smsapi.New("USERNAME", "PASSWORD")
	fmt.Println(client.SendEcoSMS("XXXXXXXXX", "Yeah, it's awesome!"))
}
