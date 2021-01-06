package segments

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	videoScale        = "-vf"
	hlsCRF            = "-crf"
	videoBitrateMap   = "-b:v:%d"
	iFrameInt         = "-g" // measured in terms of number of frames
	iFrameMin         = "-keyint_min"
	videoProfile      = "-profile:v"
	sceneCutDetection = "-sc_threshold"
	inputFile         = "-i"
	videoStreamMap    = "-c:v:%d"
	streamMap         = "-var_stream_map"
)

var (
	videoConfDefaults = map[string]interface{}{
		sceneCutDetection: "0",
		hlsCRF:            20,
		iFrameInt:         "25",
	}
)

type VideoConfig struct {
	Rend            Renditions
	VideoCodec      VideoCodecs
	ConstRateFactor int
	videoFile       string
	IframeInterval  int
	Profile         VideoProfile
}

func (v VideoConfig) isValid() error {
	if resErr := v.Rend.isValid(); resErr != nil {
		return resErr
	}

	if v.IframeInterval < 0 {
		return errors.New("Invlid IFrame Interval")
	}

	if v.ConstRateFactor > 51 || v.ConstRateFactor < 0 {
		return errors.New("Invalid Constant Rate Factor must be between 0-51")
	}

	if v.videoFile == "" {
		return errors.New("Video File has not been set.")
	}

	return nil
}

func (v *VideoConfig) SetVideoFile(fileName string) error {
	if _, e := os.Stat(fileName); e != nil {
		return errors.New("Failed to open file " + v.videoFile)
	} else if os.IsNotExist(e) {
		return errors.New(v.videoFile + "File does not exist")
	}
	v.videoFile = fileName
	return nil
}

func (v *VideoConfig) setVideoMaps(res []Resolution) []string {
	args := make([]string, 0)
	for i, r := range res {
		CodecMap := fmt.Sprintf(videoStreamMap, i)
		BitrateMap := fmt.Sprintf(videoBitrateMap, i)
		args = append(args, "-map", r.varName, CodecMap, v.VideoCodec.String(), BitrateMap, r.Bitrate)
	}
	return args
}

func (v *VideoConfig) setAudioMaps(streams int) []string {
	args := make([]string, 2*streams)
	for i := 0; i < 2*streams; i += 2 {
		args[i] = "-map"
		args[i+1] = "a:0"
	}
	return args
}

func (v *VideoConfig) SetVarStreamMap() []string {
	args := make([]string, 0)
	args = append(args, streamMap)
	stream := "\""

	for i, res := range v.Rend.Res {
		stream += fmt.Sprintf("v:%d,a:%d,name:%vpx", i, i, res.Height)
		if i != len(v.Rend.Res)-1 {
			stream += " "
		}
	}
	stream += "\""
	args = append(args, stream)
	return args
}

func (v VideoConfig) cmdArgs() []string {
	var args = make([]string, 0)

	args = append(args, inputFile, v.videoFile)

	resolutionCmd, resolutionVars := v.Rend.cmdArgs()

	args = append(args, resolutionCmd...)                  // filter_complex
	args = append(args, v.setVideoMaps(resolutionVars)...) // -map video

	args = append(args, sceneCutDetection, videoConfDefaults[sceneCutDetection].(string))

	args = append(args, hlsCRF, strconv.Itoa(v.ConstRateFactor))

	if v.IframeInterval == 0 {
		v.IframeInterval = videoConfDefaults[iFrameInt].(int)
	}

	args = append(args, iFrameInt, strconv.Itoa(v.IframeInterval))
	args = append(args, iFrameMin, videoConfDefaults[iFrameInt].(string))
	args = append(args, videoProfile, v.Profile.String())

	args = append(args, v.setAudioMaps(len(resolutionVars))...) // -map audio
	return args
}
