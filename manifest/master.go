// manifest is reponsible for generation of master manifest files
// for adaptive bitrate streaming
package manifest

import (
	"github.com/AlekSi/pointer"
	"github.com/etherlabsio/go-m3u8/m3u8"
)

func NewMaster(mediaSequence int, targetDuration int, playlistType *string) m3u8.Playlist {
	return m3u8.Playlist{
		Version:  pointer.ToInt(4),
		Master:   pointer.ToBool(true),
		Sequence: mediaSequence,
		Target:   targetDuration,
		Type:     playlistType, // Live, VOD or Event
	}
}

// func AppendToMaster(playlist *m3u8.Playlist, s []segments.StreamINF) {
// 	for _, segment := range s {
// 		playlist.AppendItem(createPlaylistItem(segment))
// 	}
// }

// func createPlaylistItem(s segments.StreamINF) *m3u8.PlaylistItem {
// 	// #EXT-X-STREAM-INF
// 	return &m3u8.PlaylistItem{
// 		Width:            &s.Width,
// 		Height:           &s.Height,
// 		Profile:          &itemProfile,
// 		Bandwidth:        s.Bandwidth,
// 		Level:            &level,
// 		AudioCodec:       &audoCodec,
// 		URI:              s.Uri,
// 		AverageBandwidth: &s.AverageBandwidth,
// 		FrameRate:        &frameRate,
// 	}
// }
