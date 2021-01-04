package segments

import (
	"errors"
	"strconv"
	"strings"
)

const (
	hlsPlaylistType     = "-hls_playlist_type"
	sequenceStartNum    = "-hls_start_number_source"
	hlsSegmentTimeSlice = "-hls_time"
)

var (
	headerConfigDefaults = map[string]int{
		sequenceStartNum:    0,
		hlsSegmentTimeSlice: 5,
	}
)

type HeaderConfig struct {
	Duration     int       // #EXT-X-TARGETDURATION - all meadia segments will be less than or equal to the Duration
	SequenceNum  int       // #EXT-X-MEDIA-SEQUENCE - the sequenceID of the first segment in the playlist
	PlaylistType MediaType // #EXT-X-PLAYLIST-TYPE - defines the mutability of the Playlist, can be either VOD or EVENT
}

func (h HeaderConfig) GetMediaType() string {
	return strings.ToLower(h.PlaylistType.String())
}

func (h HeaderConfig) isValid() error {
	if h.Duration < 0 {
		return errors.New("Duration cannot be below 0.")
	}

	if h.SequenceNum < 0 {
		return errors.New("Sequence Number must be non negative")
	}
	return nil
}

func (h HeaderConfig) cmdArgs() []string {
	args := make([]string, 0)

	args = append(args, hlsSegmentTimeSlice)
	if h.Duration == 0 {
		// set the Default args for this argument
		h.Duration, _ = headerConfigDefaults[hlsSegmentTimeSlice]
	}
	args = append(args, strconv.Itoa(h.Duration))

	args = append(args, sequenceStartNum, strconv.Itoa(h.SequenceNum))
	args = append(args, hlsPlaylistType, h.PlaylistType.String())

	return args
}
