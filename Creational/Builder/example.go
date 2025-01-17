package main

import "fmt"

func main() {
	var bldr = newNotificationBuilder()
	notification, err := bldr.
		SetTitle("New Notification").
		SetSubTitle("Sub Title").
		SetMessage("Message").
		SetImage("phot.png").
		SetPriority(5).
		SetType("alert").
		Build()

	if err != nil {
		fmt.Println("Error creating the notification:", err)
	} else {
		fmt.Printf("Notification: %+v", notification)
	}
}
