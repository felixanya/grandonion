package adapter

import "testing"

func TestMediaPlayerAdapter_Play(t *testing.T) {
	mpa := NewMediaPlayerAdapter("vlc")
	mpa.Play("Love story")
}
