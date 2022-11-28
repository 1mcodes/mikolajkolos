package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/1mcodes/mikolajkolos/internal/drawer"
	"github.com/1mcodes/mikolajkolos/internal/sender"
)

func init() {
	rand.Seed(time.Now().UnixNano() + int64(os.Getpid()))
}

func main() {

	participant1 := drawer.NewParticipant("Foo", "foo@foo.com", "Dear Santa...")
	participant2 := drawer.NewParticipant("Bar", "bar@bar.com", "Dear Santa...")
	participant3 := drawer.NewParticipant("Baz", "baz@baz.com", "Dear Santa...")

	participant1.SetBlacklist(drawer.Participants{participant2})
	participant2.SetBlacklist(drawer.Participants{participant1})

	participants := drawer.Participants{participant1, participant2, participant3}

	mailer := sender.NewMailer("smtp.gmail.com", "587", "santaclaus@gmail.com", "onlyGoodElvesKnowPassword")
	d := drawer.NewDrawer(participants, mailer, 100, false)
	d.Draw()
	for _, p := range d.P {
		err := os.WriteFile(p.Name+".yaml", []byte(p.String()), 0755)
		if err != nil {
			log.Printf("write file error: %s", err)
		}
	}
	d.Send()
}
