package obsws

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/gorilla/websocket"
)

// Open the WebSocket connection.
func connectWS(host string, port int) (*websocket.Conn, error) {
	url := fmt.Sprintf("ws://%s:%d", host, port)
	logger.Infof("connecting to %s", url)
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// Compute the auth challenge response.
func getAuth(password, salt, challenge string) string {
	sha := sha256.Sum256([]byte(password + salt))
	b64 := base64.StdEncoding.EncodeToString([]byte(sha[:]))

	sha = sha256.Sum256([]byte(b64 + challenge))
	b64 = base64.StdEncoding.EncodeToString([]byte(sha[:]))

	return b64
}

// // Remove some Events from a slice by index.
// func removeEvents(list []Event, is ...int) []Event {
// 	newSlice := []Event{}
// 	for idx, x := range list {
// 		found := false
// 		for _, i := range is {
// 			if idx == i {
// 				found = true
// 				break
// 			}
// 		}
// 		if !found {
// 			newSlice = append(newSlice, x)
// 		}
// 	}
// 	return newSlice
// }

// Remove some responses from a slice by index.
func removeResponses(list []response, is ...int) []response {
	newSlice := []response{}
	for idx, x := range list {
		found := false
		for _, i := range is {
			if idx == i {
				found = true
				break
			}
		}
		if !found {
			newSlice = append(newSlice, x)
		}
	}
	return newSlice
}
