## HLS segmenter for creating .m3u8 playlists and .ts segments

Example Usage:

```go
// create Video Config
videoConfig := segments.VideoConfig{
  Res: segments.Resolution{
    Height: 360,
    Width:  640,
  },
  ConstRateFactor: 19,
  IframeInterval:  60,
  Profile:         segments.Main, // h264 profile
  VideoFile:       "./assets/test.mp4",
}

headerConfig := segments.HeaderConfig {
  Duration:     5,
  PlaylistType: segments.Vod,
}

audioConfig := segments.AudioConfig{
  AudioCodec:   segments.Aac,
  AudioBitrate: "128k",
}

outputConfig := segments.OutputConfig{
  PlaylistFile:   "./frames2/short.m3u8",
  SegmentPattern: "./frames2/output_%d.ts",
}

playlist := segments.NewHlsConfig(headerConfig, videoConfig, audioConfig, outputConfig, os.Stdout, os.Stderr)

if err := playlist.Exec(); err != nil {
  log.Fatal("[ERR] %s\n", err)
}
```
