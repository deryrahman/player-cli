package main

import (
	"os"
	"strings"

	"github.com/deryrahman/player-cli/player"
	"github.com/deryrahman/player-cli/provider"
	"github.com/deryrahman/player-cli/provider/youtube"
)

var (
	p provider.Provider
)

func main() {
	if len(os.Args) < 2 {
		println("please provide a query")
		os.Exit(1)
	}
	query := strings.Join(os.Args[1:], " ")

	p = youtube.NewYoutubeProvider()
	audioMeta, err := p.GetAudio(query)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	println(audioMeta.Title)
	player.Play(audioMeta.URL, func(err error) {
		println(err.Error())
		os.Exit(1)
	})
}
