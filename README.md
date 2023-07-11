# go-push

![](https://img.shields.io/badge/golang-1.19-blue.svg?style=flat)

go-push.
This repository is inspired by [gorush](https://github.com/appleboy/gorush)


# Installation
```
go get github.com/coolishbee/go-push
```

# Usage

## FCM

```go
import(
    "github.com/coolishbee/go-push/fcm"
)

func main() {
    pushClient, err := fcm.NewFromAPIKey("apiKey", "Identifier")
    if err != nil {
        fmt.Println(err)
    }

    pushClient.PushTokens = []string{"token1", "token2"}
    //pushClient.FcmMessage.Notification.Title = "gameTitle010"
	//pushClient.FcmMessage.Notification.Body = "gameText010"
	pushClient.FcmMessage.Data = map[string]interface{}{
		"title": "gameTitle010",
		"body":  "gameText010",
	}

    resp := pushClient.Send()
    fmt.Printf("Success: %d\n", resp.Success)
    fmt.Printf("Failure: %d\n", resp.Failure)
    fmt.Printf("Identifier: %d\n", resp.Identifier)
}
```

## APNS

```go
import(
    "github.com/coolishbee/go-push/apns"
)

func main() {
	pushClient, err := apns.NewFromKeyFile(
		"keyFilePath.p8",
		"KeyID",
		"TeamID", "title12", "body12", false, "Identifier")

	if err != nil {
		t.Error(err)
	}

    pushClient.DeviceTokens = []string{"token1", "token2"}
    pushClient.ApnsNotification.Topic = "bundleIdentifier"
    resp := pushClient.Send()
    fmt.Printf("Success: %d\n", resp.Success)
    fmt.Printf("Failure: %d\n", resp.Failure)
    fmt.Printf("Identifier: %d\n", resp.Identifier)
}
```

