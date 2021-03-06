package segments

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	baseCmd    = "ffmpeg"
	masterConf = "-master_pl_name"
)

type HlsConfig struct {
	Header         HeaderConfig
	Output         OutputConfig
	Video          VideoConfig
	Audio          AudioConfig
	OutputFile     io.Writer
	ErrorFile      io.Writer
	MasterPlaylist string
}

func (c *HlsConfig) getArgs() []string {
	args := make([]string, 0)
	args = append(args, c.Video.cmdArgs()...)
	args = append(args, c.Audio.cmdArgs()...)
	args = append(args, c.Header.cmdArgs()...)
	args = append(args, masterConf, c.MasterPlaylist)
	args = append(args, c.Video.SetVarStreamMap()...)
	args = append(args, c.Output.cmdArgs()...)
	return args
}

func (c *HlsConfig) Exec() error {
	if e := c.isValid(); e != nil {
		return e
	}
	if e := os.MkdirAll(c.Output.BaseFolder, os.ModePerm); e != nil {
		return e
	}
	fmt.Fprintln(c.OutputFile, time.Now().UTC().Format("2006/01/02 15:04:05"), "ffmpeg", strings.Join(c.getArgs(), " "))
	cmd := exec.Command(baseCmd, c.getArgs()...)
	cmd.Stdout = c.OutputFile
	cmd.Stderr = c.ErrorFile
	if e := cmd.Run(); e != nil {
		fmt.Println(e)
		return e
	}
	return nil
}

func (c *HlsConfig) isValid() error {

	if c.MasterPlaylist == "" {
		return errors.New("Missing master playlist")
	}

	if headerErr := c.Header.isValid(); headerErr != nil {
		return headerErr
	}

	if outputErr := c.Output.isValid(); outputErr != nil {
		return outputErr
	}

	if audioErr := c.Audio.isValid(); audioErr != nil {
		return audioErr
	}

	if videoErr := c.Video.isValid(); videoErr != nil {
		return videoErr
	}

	return nil
}

func NewHlsConfig(masterPlaylist string,
	headerConfig HeaderConfig,
	videoConfig VideoConfig,
	audioConfig AudioConfig,
	outputConfig OutputConfig,
	outputFile io.Writer,
	logFile io.Writer) *HlsConfig {
	return &HlsConfig{
		MasterPlaylist: masterPlaylist,
		Video:          videoConfig,
		Header:         headerConfig,
		Audio:          audioConfig,
		Output:         outputConfig,
		ErrorFile:      logFile,
		OutputFile:     outputFile,
	}
}
