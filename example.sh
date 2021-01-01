ffmpeg -i beach.mkv \
    -vf scale=w=1280:h=720:force_original_aspect_ratio=decrease \
    -c:a aac -ar 48000 -b:a 128k -c:v h264 -profile:v main \
    -crf 20 -g 48 -keyint_min 48 -sc_threshold 0 \
    -b:v 2500k -maxrate 2675k -bufsize 3750k \
    -hls_time 4 -hls_playlist_type vod \
    -hls_segment_filename beach/720p_%03d.ts beach/720p.m3u8
