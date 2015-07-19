package tts

// Default language.
const Lang = "en"

// Speaker defines an interface for text to speech providers.
type Speaker interface {
	Say(phrase string)
}

// AsyncSpeaker wraps a Speaker for asynchronous playback.
type AsyncSpeaker struct {
	S    Speaker
	Done chan bool
}

// MakeAsync creates an AsyncSpeaker from a Speaker.
func MakeAsync(s Speaker) *AsyncSpeaker {
	return &AsyncSpeaker{s, make(chan bool)}
}

// Say wraps the underlying Speaker's Say method in a goroutine, notifying the
// caller over the Done channel (chan bool).
func (as *AsyncSpeaker) Say(phrase string) {
	go func() {
		as.S.Say(phrase)
		as.Done <- true
	}()
}
