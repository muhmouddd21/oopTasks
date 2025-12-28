package main

import (
	"fmt"
	"time"
)

// Notification system

type Notification interface {
	GetCost() float64
	NotificationInfo() string
	IsSent() bool
	Send() bool
}
type BaseNotification struct {
	IDValue     string
	Message     string
	Title       string
	IsSentValue bool
	Cost        float64
}

type NotificationService struct {
	Notifications []Notification
}

func (n *NotificationService) SendNotifications() {
	for _, item := range n.Notifications {
		if item.IsSent() == false {
			item.Send()
		}
	}
}
func (n *NotificationService) AddNotification(no Notification) {
	n.Notifications = append(n.Notifications, no)
}
func (n *NotificationService) GetAllCost() float64 {
	var result float64
	for _, item := range n.Notifications {
		result += item.GetCost()
	}
	return result
}

func (b BaseNotification) NotificationInfo() string {
	return b.IDValue + " - " + b.Title + " - " + b.Message
}
func (b BaseNotification) IsSent() bool {
	return b.IsSentValue
}

func (b *BaseNotification) Send() bool {
	fmt.Printf("sent notification of id %v \n", b.IDValue)
	b.IsSentValue = true
	return b.IsSentValue
}

type EmailNotification struct {
	BaseNotification
	EmailAddress string
}
type SMSNotification struct {
	BaseNotification
	PhoneNumber string
}
type WHatsNotification struct {
	BaseNotification
	WhatsAppNo string
}
type ScheduledNotification struct {
	Notification Notification
	SendAt       time.Time
}

func (s *ScheduledNotification) Send() bool {
	if time.Now().Before(s.SendAt) {
		return false
	}
	return s.Notification.Send() //
}
func (s *ScheduledNotification) GetCost() float64 {
	return s.Notification.GetCost()
}

func (s *ScheduledNotification) NotificationInfo() string {
	return s.Notification.NotificationInfo()
}

func (s *ScheduledNotification) IsSent() bool {
	return s.Notification.IsSent()
}

func (e EmailNotification) GetCost() float64 {

	return e.Cost * 0.5
}
func (s SMSNotification) GetCost() float64 {
	return s.Cost * 1.5
}
func (w WHatsNotification) GetCost() float64 {

	return w.Cost * 0.1
}

func main() {

	notification1 := EmailNotification{
		BaseNotification: BaseNotification{
			IDValue:     "001",
			Message:     "hello oop",
			Title:       "OOP TASK To email",
			IsSentValue: false,
			Cost:        50.00,
		},
		EmailAddress: "mahmoudabdelhaleem@gmail.com",
	}
	notification2 := SMSNotification{
		BaseNotification: BaseNotification{
			IDValue:     "002",
			Message:     "hello oop",
			Title:       "OOP TASK To phone",
			IsSentValue: false,
			Cost:        100.00,
		},
		PhoneNumber: "+20102131564648",
	}
	notification3 := WHatsNotification{
		BaseNotification: BaseNotification{
			IDValue:     "003",
			Message:     "hello oop",
			Title:       "OOP TASK To whatsapp",
			IsSentValue: false,
			Cost:        100.00,
		},
		WhatsAppNo: "+20102131564648",
	}
	sms := &SMSNotification{
		BaseNotification: BaseNotification{
			IDValue:     "002",
			Message:     "hello oop",
			Title:       "SMS scheduled",
			IsSentValue: false,
			Cost:        100.00,
		},
		PhoneNumber: "+20102131564648",
	}
	scheduledSMS := &ScheduledNotification{
		Notification: sms,
		SendAt:       time.Now().Add(5 * time.Second),
	}

	fmt.Println(notification1.NotificationInfo())
	fmt.Println(notification1.GetCost())
	fmt.Println(notification1.IsSent())
	fmt.Println(notification2.NotificationInfo())
	fmt.Println(notification2.GetCost())
	fmt.Println(notification2.IsSent())

	notifyService := NotificationService{}
	notifyService.AddNotification(&notification1)
	notifyService.AddNotification(&notification2)
	notifyService.AddNotification(&notification3)
	notifyService.AddNotification(scheduledSMS)
	// notifyService.SendNotifications()
	notifyService.GetAllCost()
	fmt.Println(notification1.NotificationInfo())
	fmt.Println(notification1.GetCost())
	fmt.Println(notification1.IsSent())
	fmt.Println(notification2.NotificationInfo())
	fmt.Println(notification2.GetCost())
	fmt.Println(notification2.IsSent())
	fmt.Println("--------------")
	fmt.Println(notifyService.GetAllCost())

	fmt.Println("=== First send (before time) ===")
	notifyService.SendNotifications()

	time.Sleep(6 * time.Second)

	fmt.Println("=== Second send (after time) ===")
	notifyService.SendNotifications()
}
