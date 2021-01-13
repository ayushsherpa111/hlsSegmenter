package segments

import (
	"errors"
	"path"
	"path/filepath"
	"strings"
)

const (
	hlsBaseURL         = "-hls_base_url"
	hlsSegmentFileName = "-hls_segment_filename"
)

type OutputConfig struct {
	PlaylistFile   string // target .m3u8 playlist file, can be relative file path
	SegmentPattern string // Pattern that is used by ffmpeg to generate segments
	BaseURL        string // Base URL for every segment
	BaseFolder     string
}

func (o OutputConfig) isValid() error {
	if filepath.Ext(o.PlaylistFile) != ".m3u8" {
		return errors.New("Incorrect Playlist File extension.")
	}
	if filepath.Ext(o.SegmentPattern) != ".ts" {
		return errors.New("Invalid Segment Pattern")
	}
	return nil
}

func (o OutputConfig) cmdArgs() []string {
	args := make([]string, 0)

	if o.BaseURL != "" {
		if !strings.HasSuffix(o.BaseURL, "/") {
			o.BaseURL += "/"
		}
		args = append(args, hlsBaseURL, o.BaseURL)
	}

	args = append(args, hlsSegmentFileName, path.Join(o.BaseFolder, o.SegmentPattern), path.Join(o.BaseFolder, o.PlaylistFile))

	return args
}
