package adapter

import "fmt"

type MediaPlayer interface {
	Play(string)
}

type AdvancedMediaPlayer interface {
	PlayVLC()
	PlayMP4()
}

// VLCPlayer 实现AdvancedMediaPlayer接口
type VLCPlayer struct {
	FileName string
}

func (vp *VLCPlayer) PlayVLC() {
	fmt.Println("VLC Player playing vlc", vp.FileName)
}

func (vp *VLCPlayer) PlayMP4() {
	// do nothing
}

// MP4Player  实现AdvancedMediaPlayer接口
type MP4Player struct {
	FileName string
}

func (mp *MP4Player) PlayVLC() {
	// do nothing
}

func (mp *MP4Player) PlayMP4() {
	fmt.Println("MP4 Player playing mp4", mp.FileName)
}

// 创建adapter
type MediaPlayerAdapter struct {
	//MediaType string
	MediaType string
}

func (mpa *MediaPlayerAdapter) Play(fileName string) {
	var amp AdvancedMediaPlayer
	switch mpa.MediaType {
	case "vlc":
		amp = &VLCPlayer{FileName: fileName}
		amp.PlayVLC()
	case "mp4":
		amp = &MP4Player{FileName: fileName}
		amp.PlayMP4()
	default:
		// do nothing
	}
}

func NewMediaPlayerAdapter(mt string) *MediaPlayerAdapter {
	return &MediaPlayerAdapter{
		MediaType: mt,
	}
}
