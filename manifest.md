# HLS manifest file breakdown

## Manifest file

### Example

```m3u8
#EXTM3U
#EXT-X-TARGETDURATION:10

#EXTINF:9.009,
http://media.example.com/first.ts
#EXTINF:9.009,
http://media.example.com/second.ts
#EXTINF:3.003,
http://media.example.com/third.ts
```

```m3u8
#EXTM3U
```

The format identifier tag

```m3u8
#EXT-X-TARGETDURATION:N
```

This tag represents the duration of each segment in seconds. Each segment will be of length <= N.

```m3u8
#EXTINF:N
https://cdn.service.com/media/segment-001.ts
```

This tag represents a segment of duration of N seconds. the link below it points to a segment on the internet.
The URI can specify any reliable protocol to transfer data.

## Master Playlist

### Example

```m3u8
#EXTM3U
#EXT-X-VERSION:3
#EXT-X-INDEPENDENT-SEGMENTS
#EXT-X-STREAM-INF:BANDWIDTH=2665726,AVERAGE-BANDWIDTH=2526299,RESOLUTION=960x540,FRAME-RATE=29.970,CODECS="avc1.640029,mp4a.40.2"
index_1.m3u8
#EXT-X-STREAM-INF:BANDWIDTH=3956044,AVERAGE-BANDWIDTH=3736264,RESOLUTION=1280x720,FRAME-RATE=29.970,CODECS="avc1.640029,mp4a.40.2"
index_2.m3u8
#EXT-X-STREAM-INF:BANDWIDTH=995315,AVERAGE-BANDWIDTH=951107,RESOLUTION=640x360,FRAME-RATE=29.970,CODECS="avc1.4D401E,mp4a.40.2"
index_3.m3u8
#EXT-X-MEDIA:TYPE=SUBTITLES,GROUP-ID="subtitles",NAME="caption_1",DEFAULT=YES,AUTOSELECT=YES,FORCED=NO,LANGUAGE="eng",URI="index_4_0.m3u8"
```

The media can be represented in a more complex way using **Master Playlist**, which is used to define the different rendition/version of the same content
for different bandwidths.

The Master Playlist is used for Adaptive Bitrate Streaming i.e. it allows the player to make decision on which resolution of video should be played based on the
clients bandwidth to ensure that the client never experiences video buffering.

## Media Segments

A media playlist is composed of several segments using the `#EXTINF` tag along with the duration of the segment.
Segments are also uniquely identified using an integer starting from 0, usually this integer is incremented for each subsequent segment optionally with a prefix.
