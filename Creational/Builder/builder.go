package main

import "fmt"

// The NotificationBuilder has fields exported as well as a few methods
// to demonstrate
type NotificationBuilder struct {
	Title    string
	SubTitle string
	Message  string
	Image    string
	Icon     string
	Priority int
	NotType  string
}

func newNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

func (nb *NotificationBuilder) SetTitle(title string) *NotificationBuilder {
	nb.Title = title

	return nb
}

func (nb *NotificationBuilder) SetSubTitle(subtitle string) *NotificationBuilder {
	nb.Title = subtitle

	return nb
}

func (nb *NotificationBuilder) SetMessage(message string) *NotificationBuilder {
	nb.Message = message

	return nb
}

func (nb *NotificationBuilder) SetImage(image string) *NotificationBuilder {
	nb.Image = image

	return nb
}

func (nb *NotificationBuilder) SetIcon(icon string) *NotificationBuilder {
	nb.Icon = icon

	return nb
}

func (nb *NotificationBuilder) SetPriority(pri int) *NotificationBuilder {
	nb.Priority = pri

	return nb
}

func (nb *NotificationBuilder) SetType(notType string) *NotificationBuilder {
	nb.NotType = notType

	return nb
}

func (nb *NotificationBuilder) Build() (Notification, error) {
	if nb.Icon != "" && nb.SubTitle == "" {
		return Notification{}, fmt.Errorf("You need specify title when use icon")
	}

	if nb.Priority > 5 {
		return Notification{}, fmt.Errorf("Priority must be 5 or less")
	}

	return Notification{
		title:    nb.Title,
		subtitle: nb.SubTitle,
		message:  nb.Message,
		image:    nb.Image,
		icon:     nb.Icon,
		priority: nb.Priority,
		notType:  nb.NotType,
	}, nil
}
