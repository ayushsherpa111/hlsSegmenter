package segments

import (
	"regexp"
)

var (
	headerDefaults = map[string]string{
		hlsPlaylistType:     "vod",
		sequenceStartNum:    "0",
		hlsSegmentTimeSlice: "5",
	}
)

func regexpMatch(pattern, selectedCodec string) bool {
	var vCodec, e = regexp.Compile(pattern)
	if e != nil {
		return false
	}
	return vCodec.MatchString(selectedCodec)
}

type Config interface {
	isValid() error
	Exec() error
}

type StreamINF struct {
	Width            int
	Height           int
	Bandwidth        int
	AverageBandwidth int
	Uri              string
}

type ExtINF struct {
	Duration float64
	URI      string
}
