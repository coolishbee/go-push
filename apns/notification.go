package apns

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/payload"
	"github.com/sideshow/apns2/token"
)

type PushClient struct {
	apnsClient       *apns2.Client
	apnsNotification *apns2.Notification
	deviceTokens     []string
}

type PushResponse struct {
	Success int
	Failure int
}

func NewFromKeyFile(
	filename string,
	keyId string,
	teamId string,
	title string,
	body string,
	production bool,
) (*PushClient, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return NewFromKeyBytes(bytes, keyId, teamId, title, body, production)
}

func NewFromKeyBytes(
	bytes []byte,
	keyId string,
	teamId string,
	title string,
	body string,
	production bool,
) (*PushClient, error) {
	authKey, err := token.AuthKeyFromBytes(bytes)
	if err != nil {
		return nil, err
	}
	token := &token.Token{
		AuthKey: authKey,
		KeyID:   keyId,
		TeamID:  teamId,
	}

	var client *apns2.Client

	if production {
		client = apns2.NewTokenClient(token).Production()
	} else {
		client = apns2.NewTokenClient(token).Development()
	}

	notification := &apns2.Notification{}

	payload := payload.NewPayload()
	payload.AlertTitle(title)
	payload.AlertBody(body)
	payload.Sound("default")
	notification.Payload = payload

	pushClient := &PushClient{
		apnsClient:       client,
		apnsNotification: notification,
	}
	return pushClient, nil
}

func (c *PushClient) Send() (resp *PushResponse) {

	resp = &PushResponse{}

	var wg sync.WaitGroup
	for _, token := range c.deviceTokens {
		wg.Add(1)
		go func(notification apns2.Notification, token string) {
			notification.DeviceToken = token

			res, err := c.apnsClient.Push(&notification)
			if err != nil {
				log.Println("There was an error", err)
				return
			}

			if res.Sent() {
				log.Println("Sent:", res.ApnsID)
				resp.Success += 1
			} else {
				fmt.Printf("Not Sent: %v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
				resp.Failure += 1
			}

			wg.Done()

		}(*c.apnsNotification, token)
	}
	wg.Wait()

	return resp
}

// func New() *Service {
// 	authKey, err := token.AuthKeyFromFile("../AuthKey.p8")
// 	if err != nil {
// 		log.Fatal("token error:", err)
// 	}
// 	token := &token.Token{
// 		AuthKey: authKey,
// 		KeyID:   "",
// 		TeamID:  "",
// 	}

// 	client := apns2.NewTokenClient(token)

// 	notification := &apns2.Notification{
// 		DeviceToken: "",
// 		Topic:       "",
// 	}

// 	payload := payload.NewPayload()
// 	payload.AlertTitle("sdk01")
// 	payload.AlertBody("redis01")
// 	payload.Sound("default")
// 	notification.Payload = payload

// 	service := &Service{
// 		APNS: client,
// 		Noti: notification,
// 	}
// 	return service
// }

// func (c *PushClient) Send() {

// 	c.apnsNotification.Expiration = time.Now()
// 	c.apnsNotification.Priority = apns2.PriorityLow

// 	res, err := c.apnsClient.Push(c.apnsNotification)
// 	if err != nil {
// 		log.Println("There was an error", err)
// 		return
// 	}

// 	if res.Sent() {
// 		log.Println("Sent:", res.ApnsID)
// 	} else {
// 		fmt.Printf("Not Sent: %v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
// 	}
// }