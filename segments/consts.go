package segments

type VideoCodecs int
type AudioCodecs int
type MediaType int
type VideoProfile int

const (
	H264 VideoCodecs = iota
	H265
)

const (
	Aac AudioCodecs = iota
	Ac3
	Eac3
	Mp3
)

const (
	Vod MediaType = iota
	Event
)

const (
	Main     VideoProfile = iota
	BaseLine              // Does not use b-frames
	High
	bitRateRegExp = "^([0-9]+)k$"
)
