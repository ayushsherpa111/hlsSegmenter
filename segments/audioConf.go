package segments

import (
	"errors"
	"strconv"
)

const (
	audioCodecOption  = "-c:a"
	audioBitrate      = "-b:a"
	audioSamplingRate = "-ar"
)

var (
	audioConfDefaults = map[string]interface{}{
		audioCodecOption:  "aac",
		audioBitrate:      "128k",
		audioSamplingRate: 48000,
	}
)

type AudioConfig struct {
	AudioCodec   AudioCodecs
	AudioBitrate string
	SamplingRate int
}

func (a AudioConfig) isValid() error {

	if a.AudioBitrate != "" && !regexpMatch(bitRateRegExp, a.AudioBitrate) {
		return errors.New("Invalid Bitrate")
	}

	if a.SamplingRate < 0 {
		return errors.New("Invalid Sampling Rate")
	}

	return nil
}

func (a AudioConfig) cmdArgs() []string {
	args := make([]string, 0)

	args = append(args, audioCodecOption, a.AudioCodec.String())

	if a.AudioBitrate == "" {
		a.AudioBitrate = audioConfDefaults[audioBitrate].(string)
	}
	args = append(args, audioBitrate, a.AudioBitrate)

	if a.SamplingRate == 0 {
		a.SamplingRate = audioConfDefaults[audioSamplingRate].(int)
	}
	args = append(args, audioSamplingRate, strconv.Itoa(a.SamplingRate))

	return args
}
