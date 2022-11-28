package drawer

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
)

type Sender interface {
	Send(to []string, msg []byte) error
	PrepareMessage(who, whom, tip string) string
}

type participant struct {
	mail, Name, tip string
	blacklist       Participants
}

type Participant struct {
	participant
	drawn *participant
}

type Participants []*Participant

type Drawer struct {
	P  Participants
	s  Sender
	t  int
	dr bool
}

func NewDrawer(p Participants, s Sender, t int, dr bool) Drawer {
	return Drawer{
		P:  p,
		s:  s,
		t:  t,
		dr: dr,
	}
}

func (d *Drawer) String() string {
	n := []string{}
	for i, v := range d.P {
		s := fmt.Sprintf("[%d] %s", i, v.Name)
		if v.drawn != nil {
			s = s + fmt.Sprintf(" drawn: %s", v.drawn.Name)
		}
		n = append(n, s)
	}
	return strings.Join(n, "\n")
}

func (d *Drawer) Draw() {
	var j, tries int
	w := make(Participants, len(d.P))
	copy(w, d.P)
	rand.Shuffle(len(d.P), func(a, b int) { d.P[a], d.P[b] = d.P[b], d.P[a] })
	for _, who := range d.P {
		whom := who
		for who.hasOnBlacklist(whom) || who.Name == whom.Name {
			rand.Shuffle(len(w), func(a, b int) { w[a], w[b] = w[b], w[a] })
			j, whom = d.draw(w)
			tries++
			if tries > d.t {
				log.Panicf("unable to generate pairs")
			}
		}
		tries = 0
		who.setDrawn(whom)
		w.removeElement(j)
	}
}

func (d *Drawer) Send() {
	if !d.dr {
		for _, who := range d.P {
			err := d.s.Send([]string{who.mail}, []byte(d.s.PrepareMessage(who.Name, who.drawn.Name, who.drawn.tip)))
			if err != nil {
				log.Printf("unable to send mail to %s: %s", who.mail, err)
			}
		}
	}
}

func NewParticipant(name, mail, tip string) *Participant {
	p := &Participant{}
	p.Name = name
	p.mail = mail
	p.tip = tip

	return p
}

func (p *Participant) SetBlacklist(s Participants) {
	p.blacklist = s
}

func (d Drawer) draw(w Participants) (int, *Participant) {
	if len(w) == 1 {
		return 0, w[0]
	}
	i := rand.Intn(len(w) - 1)
	return i, w[i]
}

func (p *Participant) setDrawn(s *Participant) {
	p.drawn = s.toParticipantPtr()
}

func (p Participant) toParticipantPtr() *participant {
	var r participant
	r.Name = p.Name
	r.mail = p.mail
	r.tip = p.tip
	r.blacklist = p.blacklist

	return &r
}

func (p *Participants) removeElement(i int) {
	c := *p
	c = append(c[:i], c[i+1:]...)
	*p = c
}

func (p *Participant) hasOnBlacklist(g *Participant) bool {
	for _, b := range p.blacklist {
		if fmt.Sprintf("%s%s", g.mail, g.Name) == fmt.Sprintf("%s%s", b.mail, b.Name) {
			return true
		}
	}
	return false
}

func (p *Participant) String() string {
	return fmt.Sprintf("name: %s\nmail: %s\ndrawn:\n  name: %s\n  mail: %s\n", p.Name, p.mail, p.drawn.Name, p.drawn.mail)
}
