// +build darwin

package atts

import "os/exec"

// Speaker is a driver using Apple's TTTS
type Speaker struct{}

// Say uses OS X builtin text to speech
func (s *Speaker) Say(phrase string) {
	exec.Command("say", phrase).Run()
}
