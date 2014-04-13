// Package main
package main
import (
    "fmt"
	"github.com/anachronistic/apns"
)

func send() {
	payload := apns.NewPayload()
	payload.Alert = "Hello, world!"
	payload.Badge = 42
	payload.Sound = "bingbong.aiff"

	pn := apns.NewPushNotification()
	pn.DeviceToken = "1111"
	pn.AddPayload(payload)
	pn.Set("foo", "bar")
	pn.Set("doctor", "who?")
	pn.Set("the_ultimate_answer", 42)

	conn := apns.NewConnection("localhost:2195")
	conn.SetTimeoutSeconds(5)

	// You can pass either base64 strings or a path to your certificate files.
	// The final parameter, rawBytes (set to false here), dictates whether the
	// first two string parameters represent the base64 data or paths.
	conn.CreateCertificate("/home/e-max/workspace/go/src/github.com/e-max/pusher/cert/certificate.pem", "/home/e-max/workspace/go/src/github.com/e-max/pusher/cert/key.pem", false)

	conn.Connect()
	defer conn.Disconnect()

	for {
		resp := conn.Send(pn)
		alert, _ := pn.PayloadString()
		fmt.Println(" Alert:", alert)
		fmt.Println("Success:", resp.Success)
		fmt.Println(" Error:", resp.Error)
	}

}
