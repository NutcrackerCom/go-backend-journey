package dllplaylist

import (
	"github.com/NutcrackerCom/go-backend-journey/doubly-linked-list/dll"
)

type DllPlaylist struct {
	Id           int
	Size         int
	CurrentMusic *dll.Node
	List         *dll.List
	Playing      bool
}

func NewPlaylistDll(l *dll.List) *DllPlaylist {
	return &DllPlaylist{List: l}
}

func (l *DllPlaylist) Add(name string) {
	added := l.List.Push_back(name)
	if l.CurrentMusic == nil {
		l.CurrentMusic = added
		l.Id = 1
	}
	l.Size++
}

func (l *DllPlaylist) Next() bool {
	if next := dll.Next(l.CurrentMusic); l.CurrentMusic != nil && next != nil {
		l.CurrentMusic = next
		l.Id++
		return true
	}
	return false
}

func (l *DllPlaylist) Prev() bool {
	if prev := dll.Prev(l.CurrentMusic); l.CurrentMusic != nil && prev != nil {
		l.CurrentMusic = prev
		l.Id--
		return true
	}
	return false
}

func (l *DllPlaylist) GetCurrentMusic() string {
	if l.CurrentMusic != nil {
		return dll.GetNode(l.CurrentMusic)
	}
	return "NO MUSIC"
}

func (l *DllPlaylist) GetAllMusic() []string {
	var allMusic []string
	for node := l.CurrentMusic; node != nil; node = dll.Next(node) {
		allMusic = append(allMusic, dll.GetNode(node))
	}
	return allMusic
}

func (l *DllPlaylist) AddNext(name string) {
	l.List.Insert(l.Id, name)
	//l.Next()

}

func (l *DllPlaylist) Stop() {
	l.Playing = false
}

func (l *DllPlaylist) Play() {
	l.Playing = true
	/*
		to do
		Тут надо будет смотреть канал и если придет Stop
		то прервать цикл и остановиться на том треке
	*/
}
