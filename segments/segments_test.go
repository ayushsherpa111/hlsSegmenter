package segments

import (
	"fmt"
	"strings"
	"testing"
)

type subConfig interface {
	cmdArgs() []string
	isValid() error
}

type testCase struct {
	config subConfig
	result string
}

func TestHeaderConfig(t *testing.T) {
	testCase := []testCase{
		{
			config: HeaderConfig{
				Duration:     6,
				PlaylistType: Vod,
			},
			result: "-hls_time 6 -hls_start_number_source 0 -hls_playlist_type vod",
		},
		{
			config: HeaderConfig{
				PlaylistType: Event,
			},
			result: "-hls_time 5 -hls_start_number_source 0 -hls_playlist_type event",
		},
		{
			config: HeaderConfig{
				SequenceNum:  10,
				PlaylistType: Vod,
			},
			result: "-hls_time 5 -hls_start_number_source 10 -hls_playlist_type vod",
		},
	}

	runConfigTests(testCase, t)
}

func TestOutputConfig(t *testing.T) {
	testCase := []testCase{
		{
			config: &OutputConfig{
				BaseURL:        "https://localhost:8080/",
				PlaylistFile:   "output.m3u8",
				SegmentPattern: "segment_%t.ts",
			},
			result: "-hls_base_url https://localhost:8080/ -hls_segment_filename segment_%t.ts output.m3u8",
		},
		{
			config: &OutputConfig{
				PlaylistFile:   "output.m3u8",
				SegmentPattern: "./%v/test_%03d.ts",
			},
			result: "-hls_segment_filename ./%v/test_%03d.ts output.m3u8",
		},
	}

	runConfigTests(testCase, t)
}

func TestAudioConfig(t *testing.T) {
	var testCases = []testCase{
		{
			config: AudioConfig{
				AudioCodec:   Eac3,
				AudioBitrate: "128k",
				SamplingRate: 44000,
			},
			result: "-c:a eac3 -b:a 128k -ar 44000",
		},
		{
			config: AudioConfig{},
			result: "-c:a aac -b:a 128k -ar 44200",
		},
		{
			config: AudioConfig{
				AudioCodec:   Mp3,
				AudioBitrate: "130k",
			},
			result: "-c:a mp3 -b:a 130k -ar 44200",
		},
	}

	runConfigTests(testCases, t)
}

func TestVideoConfig(t *testing.T) {
	var testCases = []testCase{
		{
			config: VideoConfig{
				Rend: Renditions{[]Resolution{
					{Width: 1080, Height: 720, Bitrate: "3000k"},
					{Width: 960, Height: 540, Bitrate: "2000k"},
				}},
				VideoCodec:      H264,
				ConstRateFactor: 23,
				videoFile:       "../assets/short.mp4",
				IframeInterval:  29 * 2,
				Profile:         Main,
			},
		},
	}
	// runConfigTests(testCases, t)
	fmt.Println(testCases[0].config.cmdArgs())
}

func TestRenditions(t *testing.T) {
	renditions := Renditions{
		Res: []Resolution{
			{Width: 1080, Height: 720, Bitrate: "3000k"},
			{Width: 960, Height: 540, Bitrate: "2000k"},
			{Width: 768, Height: 432, Bitrate: "1100k"},
		},
	}
	fmt.Println(renditions.cmdArgs())
}

func runConfigTests(testCases []testCase, t *testing.T) {
	for _, test := range testCases {
		if e := test.config.isValid(); e != nil {
			t.Errorf("Invalid Config: %s ", e)
		}
		gotResult := strings.Join(test.config.cmdArgs(), " ")
		if len(gotResult) != len(test.result) {
			t.Errorf("Expected: %s, Got: %s", test.result, gotResult)
		}

		if stringAscii(gotResult) != stringAscii(test.result) {
			t.Errorf("Mismatch output, Expected: %s, Got: %s", test.result, gotResult)
		}
	}
}

func stringAscii(v string) int {
	var sum int
	for _, v := range v {
		sum += int(v)
	}
	return sum
}
