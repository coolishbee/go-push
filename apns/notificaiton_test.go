package apns

import (
	"fmt"
	"testing"
)

var deviceTokens = []string{"123", "11aa01229f15f0f0c52029d8cf8cd0aeaf2365fe4cebc4af26cd6d76b7919ef7", "11aa01229f15f0f0c52029d8cf8cd0aeaf235fe4cebc4af26cd6d76b7919ef7", "11aa01229f15f0f0c52029d8cf8cd0aeaf265fe4cebc4af26cd6d76b7919ef7", "11aa01229f15f0f0c52029d8cf8cd0aeaf235fe4cebc4af26cd6d76b7919ef7", "11aa01229f15f0f0c52029d8cf8cd0aeaf265fe4cebc4af26cd6d76b7919ef7", "11aa01229f15f0f0c52029d8cf8cd0aeaf235fe4cebc4af26cd6d76b7919ef7", "11aa01229f15f0f0c52029d8cf8cd0aeaf365fe4cebc4af26cd6d76b7919ef7", "11aa01229f15f0f0c52029d8cf8cd0aeaf23fe4cebc4af26cd6d76b7919ef7", "11aa01229f15f0f0c52029d8cf8cd0aeaf2365fe4cebc4af26cd6d76b7919ef7", "11aa01229f15f0f0c52029d8cf8cd0aeaf2365fe4cebc4af26cd6d76b7919ef7", "11aa01229f15f0f0c5209d8c8cd0aeaf2365fe4cebc4af26cd6d76b7919ef7", "11aa01229f15f0f0c52029d8cf8cd0aaf2365fe4cebc4af26cd6d76b7919ef7", "11aa01229f15f0f0c52029df8cd0aeaf2365fe4cebc4af26cd6d76b7919ef7", "11aa01229f15f0f0c5029d8cf8cd0aeaf2365fe4cebc4af26cd6d76b7919ef7", "11aa01229f15f0f0c5029d8cf8cd0aeaf2365fe4cebc4af26cd6d76b7919ef7", "11aa01229f15f0f0c2029d8cf8cd0aeaf265fe4cebc4af26cd6d76b7919ef7", "11aa01229f15f0f0c52029d8cf8cd0aea235fe4cebc4af26cd6d76b7919ef7", "11aa01229f15f0f0c52029d8cf8cd0aef2365fe4cebc4af26cd6d76b7919ef7", "11aa01229f15f0f0c52029d8cf8cd0eaf2365fe4cebc4af26cd6d76b7919ef7", "11aa01229f15f0f0c52029d8cf8cd0aea2365fe4cebc4af26cd6d76b7919ef7", "11aa01229f15f0f0c52029d8cf8cd0aeaf265fe4cebc4af26cd6d76b7919ef7", "11aa01229f15f0f0c52029d8cf8cd0aeaf265fe4cebc4af26cd6d76b7919ef7", "11aa01229f15f0f0c52029d8cf8cd0aeaf265fe4cebc4af26cd6db7919ef7", "11aa01229f15f0f0c52029d8cf8cd0aeaf2365fe4cebc4af26cd66b7919ef7", "14f6c9f0aa26ee75c185b81d934b03361764f210ec1cf39677f3d04279bcf766"}

var keyFilePath = "../AuthKey.p8"

func TestAPNS(t *testing.T) {
	pushClient, err := NewFromKeyFile(
		keyFilePath,
		"",
		"", "title", "body", false)

	if err != nil {
		t.Error(err)
	}
	pushClient.deviceTokens = deviceTokens
	pushClient.apnsNotification.Topic = "bundleid"
	resp := pushClient.Send()
	fmt.Printf("Success: %d\n", resp.Success)
	fmt.Printf("Failure: %d\n", resp.Failure)
}
