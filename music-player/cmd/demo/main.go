package main

import (
	"fmt"
	"log"
	"os"

	"github.com/NutcrackerCom/go-backend-journey/doubly-linked-list/dll"
	"github.com/NutcrackerCom/go-backend-journey/music-player/internal/player"
	dllplaylist "github.com/NutcrackerCom/go-backend-journey/music-player/internal/playlist/dll-playlist"
	sliceplaylist "github.com/NutcrackerCom/go-backend-journey/music-player/internal/playlist/slice-playlist"
)

func main() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	var list dll.List
	listDll := dllplaylist.NewPlaylistDll(&list)
	playDll := player.NewMusicPlayer(listDll)
	playDll.Add("one")
	playDll.Add("two")
	playDll.AddNext("^^")
	playDll.GetAllMusic()

	fmt.Println("+++++++++++++++++++++")
	listSlice := sliceplaylist.NewPlaylistSlice()
	playSlice := player.NewMusicPlayer(listSlice)
	playSlice.Add("one")
	playSlice.Add("two")
	playSlice.AddNext("^^")
	playSlice.GetAllMusic()
}
