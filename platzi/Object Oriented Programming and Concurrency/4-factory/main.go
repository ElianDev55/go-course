package main

import "fmt"

type INotification interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}


type SMSNotification struct {

}

func (n *SMSNotification) SendNotification() {
  fmt.Println("Sending SMS notification...")
}

func (n *SMSNotification) GetSender() ISender {
  return &SMSNotificationSender{}
}

type SMSNotificationSender struct {

}

func (s *SMSNotificationSender) GetSenderMethod() string {
  return "SMS"
}

func (s *SMSNotificationSender) GetSenderChannel() string {
  return "SMS Channel"
}

