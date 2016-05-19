# FFMPEG
A simple package for using FFMPEG

##Open a video file
```go
package mypackage

import (
	"ffmpeg"
	"fmt"
)

func MyMovie() {
	movie := ffmpeg.Movie{}
	var path string = "/Users/leon/Downloads/4.mp4"
	err := movie.Initialize(&path)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(movie)
	}
}
```

##Transcode a video file
```go
package mypackage

import (
	"ffmpeg"
	"fmt"
	"path/filepath"
)

func MyTranscoder() {
	path_str, _ := filepath.Abs("movies/awesome_widescreen.mov")
	movie := ffmpeg.Movie{}
	movie.Initialize(&path_str)
	out_path_str, _ := filepath.Abs("movies/output/awesome_widescreen.mp4")
	transcoder := ffmpeg.Transcoder{Movie: movie, OutputFile: out_path_str, RawOption: ffmpeg.EncodingOption{VideoCodec: "h264"}}
	transcoder.Run()
	fmt.Println(transcoder.Movie)
}

```

##Append a video file to another
```go
package mypackage

import (
	"ffmpeg"
	"fmt"
	"path/filepath"
)

func MyAppend() {
	path_str, _ := filepath.Abs("movies/1.mp4")
	advert_path_str, _ := filepath.Abs("movies/awesome_widescreen.mov")
	movie := ffmpeg.Movie{}
	movie.Initialize(&path_str)
	out_path_str, _ := filepath.Abs("movies/output/awesome_widescreen_1.mp4")
	transcoder := ffmpeg.Transcoder{}
	transcoder.Initialize(&movie, &out_path_str, &ffmpeg.EncodingOption{VideoCodec: "h264", AppendMovie: advert_path_str}, &ffmpeg.TranscoderOption{})
	transcoder.Append()
}

```

##Prepend a video file to another
```go
package mypackage

import (
	"ffmpeg"
	"fmt"
	"path/filepath"
)

func MyPrepend() {
	path_str, _ := filepath.Abs("movies/1.mp4")
	advert_path_str, _ := filepath.Abs("movies/awesome_widescreen.mov")
	movie := ffmpeg.Movie{}
	movie.Initialize(&path_str)
	out_path_str, _ := filepath.Abs("movies/output/awesome_widescreen_2.mp4")
	transcoder := ffmpeg.Transcoder{}
	transcoder.Initialize(&movie, &out_path_str, &ffmpeg.EncodingOption{VideoCodec: "h264", AppendMovie: advert_path_str}, &ffmpeg.TranscoderOption{})
	transcoder.Prepend()
}

```

###Args for ffmpeg
```json
Aspect                    -aspect
VideoCodec                -vcodec
FrameRate                 -r
Resolution                -s
VideoBitrate              -b:v
AudioCodec                -acodec
AudioBitrate              -b:a
AudioSampleRate           -ar
AudioChannels             -ac
VideoMaxBitrate           -maxrate
VideoMinBitrate           -minrate
BufferSize                -bufsize
VideoBitrateTolerance     -bt
Threads                   -threads
Target                    -target
Duration                  -t
VideoPreset               -vpre
AudioPreset               -apre
FilePreset                -fpre
KeyframeInterval          -g
SeekTime                  -ss
Quality                   -q:v
Vframes                   -vframes
X264Vprofile              -vprofile
X264VPreset               -preset
Watermark                 -i
```