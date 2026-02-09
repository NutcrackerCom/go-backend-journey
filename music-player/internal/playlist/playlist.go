package playlist

type Playlist interface {
	Add(name string)
	AddNext(name string)
	Play()
	Stop()
	Next() bool
	Prev() bool
	GetCurrentMusic() string
	GetAllMusic() []string
}
