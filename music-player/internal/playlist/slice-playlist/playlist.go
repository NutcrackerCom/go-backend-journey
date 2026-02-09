package sliceplaylist

type SlicePlaylist struct {
	Id           int
	Size         int
	CurrentMusic string
	List         []string
	Playing      bool
}

func NewPlaylistSlice() *SlicePlaylist {
	return &SlicePlaylist{}
}

func (l *SlicePlaylist) Add(name string) {
	l.List = append(l.List, name)
	if l.CurrentMusic == "" {
		l.CurrentMusic = name
	}
	l.Size++
}

func (l *SlicePlaylist) Next() bool {
	if l.Id < l.Size {
		l.CurrentMusic = l.List[l.Id+1]
		l.Id++
		return true
	}
	return false
}

func (l *SlicePlaylist) Prev() bool {
	if l.Id > 0 {
		l.CurrentMusic = l.List[l.Id-1]
		l.Id--
		return true
	}
	return false
}

func (l *SlicePlaylist) GetCurrentMusic() string {
	return l.List[l.Id]
}

func (l *SlicePlaylist) GetAllMusic() []string {
	return l.List
}

func (l *SlicePlaylist) AddNext(name string) {
	l.List = append(l.List, name)
	copy(l.List[l.Id+1:], l.List[l.Id:])
	l.List[l.Id+1] = name
}

func (l *SlicePlaylist) Stop() {
	l.Playing = false
}

func (l *SlicePlaylist) Play() {
	l.Playing = true
	/*
		to do
		Тут надо будет смотреть канал и если придет Stop
		то прервать цикл и остановиться на том треке
	*/
}
