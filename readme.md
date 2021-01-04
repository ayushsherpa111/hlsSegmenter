## HLS segmenter for creating .m3u8 playlists and .ts segments

Example Usage:

```go
// create Video Config
videoConfig := segments.VideoConfig{
    Rend: segments.Renditions{
        Res: []segments.Resolution{
            {Width: 1080, Height: 720, Bitrate: "3000k"},
            {Width: 960, Height: 540, Bitrate: "2000k"},
        }},
    VideoCodec:      segments.H264,
    ConstRateFactor: 21,
    IframeInterval:  24 / 2,
    Profile:         segments.High,
}
videoConfig.SetVideoFile("./assets/bunny_slice.mp4")

audioConfig := segments.AudioConfig{
    AudioCodec:   segments.Aac,
    AudioBitrate: "128k",
    SamplingRate: 48000,
}

headerConfig := segments.HeaderConfig{
    Duration:     5,
    PlaylistType: segments.Vod,
    SequenceNum:  0,
}

outputConfig := segments.OutputConfig{
    BaseURL:        "https://localhost:8000/",
    PlaylistFile:   "./frames2/%v/stream.m3u8",
    SegmentPattern: "./frames2/%v/output_%d.ts",
}

playlistConf := segments.NewHlsConfig("master.m3u8", headerConfig, videoConfig, audioConfig, outputConfig, os.Stdout, os.Stderr)

if err := playlist.Exec(); err != nil {
  log.Fatal("[ERR] %s\n", err)
}
```
