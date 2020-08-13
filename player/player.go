package player

import (
	"fmt"
	"io"
	"os/exec"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func Play(url string, callback func(err error)) {
	r, w := io.Pipe()
	defer w.Close()

	cmd := exec.Command("ffmpeg", "-i", url, "-f", "mp3", "-")
	cmd.Stdout = w

	go func() {
		streamer, format, err := mp3.Decode(r)
		if err != nil {
			callback(err)
			return
		}
		defer streamer.Close()

		done := make(chan bool)
		speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
		speaker.Play(beep.Seq(streamer, beep.Callback(func() {
			done <- true
		})))

		for {
			select {
			case <-done:
				return
			case <-time.After(time.Second):
				speaker.Lock()
				fmt.Printf("\r%s", format.SampleRate.D(streamer.Position()).Round(time.Second))
				speaker.Unlock()
			}
		}
	}()

	if err := cmd.Run(); err != nil {
		callback(err)
		return
	}
}
