package fcm

import (
	"fmt"
	"log"
	"sync"

	"github.com/appleboy/go-fcm"
)

type PushClient struct {
	FcmClient  *fcm.Client
	FcmMessage *fcm.Message
	PushTokens []string
	Identifier int
}

type PushResponse struct {
	Success    int
	Failure    int
	Identifier int
}

func NewFromAPIKey(apiKey string, id int) (*PushClient, error) {
	msg := &fcm.Message{
		Notification: &fcm.Notification{},
	}

	client, err := fcm.NewClient(apiKey)
	if err != nil {
		return nil, err
	}

	pushClient := &PushClient{
		FcmClient:  client,
		FcmMessage: msg,
		Identifier: id,
	}
	return pushClient, nil
}

func (c *PushClient) Send() (resp *PushResponse) {

	resp = &PushResponse{}
	resp.Identifier = c.Identifier

	count := 1000
	var j int
	var wg sync.WaitGroup

	for i := 0; i < len(c.PushTokens); i += count {
		j += count
		if j > len(c.PushTokens) {
			j = len(c.PushTokens)
		}
		//fmt.Printf("i : %d j : %d\n", i, j)
		//fmt.Println(Tokens[i:j])

		wg.Add(1)
		go func(fcmMessage fcm.Message, token []string) {
			//fmt.Println(token)
			fcmMessage.RegistrationIDs = token
			//fcmMessage.To = token
			res, err := c.FcmClient.Send(&fcmMessage)
			if err != nil {
				log.Fatalln(err)
			}
			resp.Success += res.Success
			resp.Failure += res.Failure

			for _, result := range res.Results {

				if result.Error != nil {
					if !result.Unregistered() {
						log.Println("Unregistered")
					}
					fmt.Println(result.Error)
					continue
				}
			}
			//fmt.Println("Succeeded: ", res.MulticastID)

			wg.Done()

		}(*c.FcmMessage, c.PushTokens[i:j])
	}

	wg.Wait()

	return resp
}

// for _, token := range Tokens {
// 	wg.Add(1)
// 	go func(fcmMessage fcm.Message, token string) {
// 		//fcmMessage.RegistrationIDs =
// 		fcmMessage.To = token
// 		res, err := c.fcmClient.Send(&fcmMessage)
// 		if err != nil {
// 			log.Fatalln(err)
// 		}

// 		for _, result := range res.Results {

// 			if result.Error != nil {
// 				if !result.Unregistered() {
// 					log.Println("Unregistered")
// 				}
// 				log.Println(result.Error)
// 				continue
// 			}
// 			log.Println("Succeeded")
// 		}

// 		wg.Done()

// 	}(*c.fcmMessage, token)
// }

// func (c *PushClient) Send() {

// 	response, err := c.fcmClient.Send(c.fcmMessage)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	for _, result := range response.Results {

// 		if result.Error != nil {
// 			if !result.Unregistered() {
// 				log.Println("Unregistered")
// 			}
// 			log.Println(result.Error)
// 			continue
// 		}
// 		log.Println("Succeeded")
// 	}

// 	//log.Printf("%#v\n", response)
// }
