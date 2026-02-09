package player

import (
	"fmt"

	"github.com/NutcrackerCom/go-backend-journey/music-player/internal/playlist"
)

type MusicStatus struct {
	Playing bool
}

type MusicPlayer struct {
	Player playlist.Playlist
	Status MusicStatus
}

func NewMusicPlayer(l playlist.Playlist) *MusicPlayer {
	return &MusicPlayer{Player: l}
}

func (mp *MusicPlayer) Add(name string) {
	mp.Player.Add(name)
}

func (mp *MusicPlayer) Next() {
	mp.Player.Next()
}

func (mp *MusicPlayer) Prev() {
	mp.Player.Prev()
}

func (mp *MusicPlayer) AddNext(name string) {
	mp.Player.AddNext(name)
}

func (mp *MusicPlayer) Play() {
	mp.Status.Playing = true
	mp.Player.Play()
}

func (mp *MusicPlayer) Stop() {
	mp.Status.Playing = false
	mp.Player.Stop()
}

func (mp *MusicPlayer) GetCurrentMusic() string {
	return mp.Player.GetCurrentMusic()
}

func (mp *MusicPlayer) GetAllMusic() {
	for _, music := range mp.Player.GetAllMusic() {
		fmt.Println(music)
	}
}
