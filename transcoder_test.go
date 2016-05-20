package ffmpeg

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestTranscoder(t *testing.T) {
	path_str, _ := filepath.Abs("movies/awesome_movie.mov")
	movie := Movie{}
	movie.Initialize(&path_str)
	out_path_str, _ := filepath.Abs("movies/output/awesome_movie.mp4")
	system(fmt.Sprintf("mkdir -p %s", strings.Replace(out_path_str, "/awesome_movie.mp4", "", -1)))
	transcoder := Transcoder{Movie: movie, OutputFile: out_path_str, RawOption: EncodingOption{VideoCodec: "h264", Threads: "1"}}
	transcoder.Run()
	fmt.Println(transcoder.Movie)
}

func TestAppend(t *testing.T) {
	path_str, _ := filepath.Abs("movies/1.mp4")
	advert_path_str, _ := filepath.Abs("movies/awesome_movie.mov")
	movie := Movie{}
	movie.Initialize(&path_str)
	out_path_str, _ := filepath.Abs("movies/output/awesome_movie_1.mp4")
	transcoder := Transcoder{}
	transcoder.Initialize(&movie, &out_path_str, &EncodingOption{VideoCodec: "h264", AppendMovie: advert_path_str, Threads: "1"}, &TranscoderOption{})
	go transcoder.Append()
	time.Sleep(time.Second)
	fmt.Println(transcoder.Process.Pid)
}

func TestPrepend(t *testing.T) {
	path_str, _ := filepath.Abs("movies/1.mp4")
	advert_path_str, _ := filepath.Abs("movies/awesome_movie.mov")
	movie := Movie{}
	movie.Initialize(&path_str)
	out_path_str, _ := filepath.Abs("movies/output/awesome_movie_2.mp4")
	transcoder := Transcoder{}
	transcoder.Initialize(&movie, &out_path_str, &EncodingOption{VideoCodec: "h264", AppendMovie: advert_path_str, Threads: "1"}, &TranscoderOption{})
	transcoder.Prepend()
}
