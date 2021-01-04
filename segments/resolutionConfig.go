package segments

import (
	"errors"
	"fmt"
	"strings"
)

const (
	filterComplex = "-filter_complex"
	videoRes      = "[v:0]split=%d"
	scaleConf     = "%sscale=%dx%d%s"
	tempVarConf   = "[v%s%03d]"
)

type Resolution struct {
	Width   int
	Height  int
	Bitrate string
	varName string
}

func (r *Resolution) setVarname() {
	r.varName = fmt.Sprintf(tempVarConf, "res", r.Height)
}

type Renditions struct {
	Res []Resolution
}

func (r Renditions) isValid() error {
	if len(r.Res) < 1 {
		return errors.New("Invalid Resolutions.")
	}

	for _, v := range r.Res {
		if !regexpMatch(bitRateRegExp, v.Bitrate) {
			return errors.New("Invalid Bitrate")
		}
	}

	return nil
}

func (r *Renditions) cmdArgs() ([]string, []Resolution) {
	args := make([]string, 0)
	args = append(args, filterComplex)
	resSplit := fmt.Sprintf(videoRes, len(r.Res))
	scaleVar := make([]string, len(r.Res))
	for i := range r.Res {
		tempVars := fmt.Sprintf(tempVarConf, "temp", i)
		r.Res[i].setVarname()
		scaleVar[i] = fmt.Sprintf(scaleConf, tempVars, r.Res[i].Width, r.Res[i].Height, r.Res[i].varName)
		resSplit += tempVars
	}
	resSplit += ";" + strings.Join(scaleVar, ";")
	args = append(args, resSplit)
	return args, r.Res
}
