package main

import "fmt"

func main() {
	//Notificaciones -> SMS, Push notification, Email
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")

	smsFactory.SendNotification()
	emailFactory.SendNotification()

	getMethod(smsFactory)
	getMethod(emailFactory)
}

type NotificationFactory interface {
	SendNotification()
	GetSender() Sender
}

type Sender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type SMSNotification struct {
}

func (SMSNotification) SendNotification() {
	fmt.Println("Sending notification via SMS")
}

func (SMSNotification) GetSender() Sender {
	return SMSNotificationSender{}
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

type EmailNotification struct {
}

func (n EmailNotification) SendNotification() {
	fmt.Println("Sending notification via EMAIL")
}

func (n EmailNotification) GetSender() Sender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "EMAIL"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

func getNotificationFactory(notificationType string) (NotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}

	if notificationType == "Email" {
		return &EmailNotification{}, nil

	}

	return nil, fmt.Errorf("No notification type, %s", "Terible")
}

func sendNotification(factory NotificationFactory) {
	factory.SendNotification()
}

func getMethod(factory NotificationFactory) {
	fmt.Println(factory.GetSender().GetSenderMethod())
}
