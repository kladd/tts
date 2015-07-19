// +build darwin

package atts

import "os/exec"

// Speaker implements tts.Speaker using Apple's TTS
type Speaker struct{}

// Say uses OS X's built-in text to speech
func (s *Speaker) Say(phrase string) {
	exec.Command("say", phrase).Run()
}
