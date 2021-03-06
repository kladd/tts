# tts - Text to speech in Go

A simple text to speech library for Go applications.

### Example

The example below uses the Google text to speech driver provided by the included [gtts](./engines/gtts) package. See the [Engines](#engines) section for available text to speech engines.

Playback can be done asynchronously if you choose as shown below. Currently language is configured by changing the constant in `tts.go`.

```go
import (
	"github.com/kladd/tts"
	"github.com/kladd/tts/engines/gtts"
)

func main() {
	speaker := new(gtts.Speaker)
	speaker.Say("Hello, World!")

	aspeaker := tts.MakeAsync(speaker)
	aspeaker.Say("Hello, asynchronous World!")
	<-aspeaker.Done
}

```

### Engines

* Google - [gtts](./engines/gtts)
* Apple (Darwin only) - [atts](./engines/atts)

### Caveats

Currently audio playback uses a shell command. On Darwin hosts this is `afplay`, on Linux `mpg123`. No support for Windows right now.

### Installation

```bash
go get github.com/kladd/tts
```

### License

See [LICENSE](./LICENSE)

### Contributing

Appreciated. Fork, branch, commit, push, pull request.
