package main

import "github.com/wenerme/astgo/agi"

func main() {
	handler := (func(session *agi.Session) {
		client := session.Client()
		client.Answer()
		client.StreamFile("tt-monkeys", "#")
		client.Exec("Dial", "PJSIP/PHONE_A") // PHONE_A must be defined in pjsip.confgit
		client.SetVariable("AGISTATUS", "SUCCESS")
		client.Hangup()
	})

	err := agi.Listen(":4573", handler)
	if err != nil {
		panic(err)
	}
}
