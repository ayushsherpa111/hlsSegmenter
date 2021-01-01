package segments

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	videoScale        = "-vf"
	videoResolution   = "scale=w=%v:h=%v:force_original_aspect_ratio=decrease"
	hlsCRF            = "-crf"
	videoCodecOption  = "-c:v"
	videoBitrate      = "-b:v"
	videoProfile      = "-profile:v"
	iFrameInt         = "-g" // measured in terms of number of frames
	iFrameMin         = "-keyint_min"
	sceneCutDetection = "-sc_threshold"
	inputFile         = "-i"
)

var (
	videoConfDefaults = map[string]interface{}{
		hlsCRF:       30,
		videoBitrate: "3000k",
		iFrameInt:    25,
	}
)

type Resolution struct {
	Width  int
	Height int
}

func (r Resolution) isValid() error {

	if r.Width < 0 {
		return errors.New("Invalid Width.")
	}

	if r.Height < 0 {
		return errors.New("Invalid Height.")
	}

	return nil
}

type VideoConfig struct {
	Res             Resolution
	VideoCodec      VideoCodecs
	VideoBitrate    string
	Profile         VideoProfile
	ConstRateFactor int
	VideoFile       string
	IframeInterval  int
}

func (v VideoConfig) isValid() error {
	if resErr := v.Res.isValid(); resErr != nil {
		return resErr
	}

	if v.IframeInterval < 0 {
		return errors.New("Invlid IFrame Interval")
	}

	if v.VideoBitrate != "" && !regexpMatch(bitRateRegExp, v.VideoBitrate) {
		return errors.New("Invalid Bitrate. Expected Format [Nk] where N = Decimal Number ")
	}

	if v.ConstRateFactor > 51 || v.ConstRateFactor < 0 {
		return errors.New("Invalid Constant Rate Factor must be between 0-51")
	}

	if _, e := os.Stat(v.VideoFile); e != nil {
		return errors.New("Failed to open file " + v.VideoFile)
	} else if os.IsNotExist(e) {
		return errors.New(v.VideoFile + "File does not exist")
	}

	return nil
}

func (v VideoConfig) cmdArgs() []string {
	var args = make([]string, 0)

	args = append(args, inputFile, v.VideoFile)

	args = append(args, videoScale, fmt.Sprintf(videoResolution, v.Res.Width, v.Res.Height))

	args = append(args, videoCodecOption, v.VideoCodec.String())

	args = append(args, sceneCutDetection, "0")
	if v.ConstRateFactor == 0 {
		v.ConstRateFactor = videoConfDefaults[hlsCRF].(int)
	}
	args = append(args, hlsCRF, strconv.Itoa(v.ConstRateFactor))

	if v.VideoBitrate == "" {
		v.VideoBitrate = videoConfDefaults[videoBitrate].(string)
	}
	args = append(args, videoBitrate, v.VideoBitrate)

	args = append(args, videoProfile, v.Profile.String())

	if v.IframeInterval == 0 {
		v.IframeInterval = videoConfDefaults[iFrameInt].(int)
	}
	args = append(args, iFrameInt, strconv.Itoa(v.IframeInterval), iFrameMin, "25")

	return args
}