#!/usr/bin/env bash

ffmpeg -hide_banner -i $1 \
    -filter_complex "[v:0]split=2[vtemp001][vtemp002];[vtemp001]scale=960x540[vlowres];[vtemp002]scale=1280x760[vhighres]" \
    -map [vlowres] -c:v:0 libx264 -b:v:0 2000k \
    -map [vhighres] -c:v:1 libx264 -b:v:1 6000k \
    -map a:0 -map a:0 -c:a aac -b:a 128k -ac 2 \
    -profile:v main -crf 20 \
    -g 30 \
    -sc_threshold 0 \
    -f hls \
    -hls_time 5 \
    -hls_start_number_source 0 \
    -var_stream_map "v:0,a:0,name:540 v:1,a:1,name:760" \
    -hls_base_url "http://localhost:8000/" \
    -hls_playlist_type vod \
    -hls_segment_filename "../frames2/%v/output_%d.ts" \
    -master_pl_name master.m3u8 \
    ../frames2/%v/stream.m3u8
