package ffmpeg

import (
	"fmt"
	"math"
	"os"
	// "strings"
)

var transcodingQuery map[int]Transcoder

type Transcoder struct {
	Process          *os.Process
	Movie            Movie
	AppendMovie      Movie
	OutputFile       string
	RawOption        EncodingOption
	TranscoderOption TranscoderOption
}

func (this *Transcoder) Initialize(movie *Movie, outputFile *string, option *EncodingOption, transcoderOption *TranscoderOption) {
	this.Movie = *movie
	if option.AppendMovie != "" {
		this.AppendMovie = Movie{}
		this.AppendMovie.Initialize(&option.AppendMovie)
	}
	this.OutputFile = *outputFile
	this.RawOption = *option
	this.TranscoderOption = *transcoderOption
	this.applyTranscoderOption()
}
func (this *Transcoder) Run() {
	this.transcodeMovie()
}

func (this *Transcoder) Append() {
	this.transcode(&this.Movie.Path, &this.Movie.TsPath)
	ffmpegBinary := FfmpegBinary()
	fmt.Println(ffmpegBinary)
	this.transcode(&this.AppendMovie.Path, &this.AppendMovie.TsPath)
	command := fmt.Sprintf("%s -y -i \"concat:%s|%s\" %s %s", ffmpegBinary, this.Movie.TsPath, this.AppendMovie.TsPath, this.RawOption.ToString(), this.OutputFile)
	systemWithoutReturn(this, command)
	// fmt.Println(strings.Replace(system(command), "\n", "", -1))
	systemWithoutReturn(this, "rm -rf "+this.Movie.TsPath)
	systemWithoutReturn(this, "rm -rf "+this.AppendMovie.TsPath)
}

func (this *Transcoder) Prepend() {
	this.transcode(&this.Movie.Path, &this.Movie.TsPath)
	ffmpegBinary := FfmpegBinary()
	fmt.Println(ffmpegBinary)
	this.transcode(&this.AppendMovie.Path, &this.AppendMovie.TsPath)
	command := fmt.Sprintf("%s -y -i \"concat:%s|%s\" %s %s", ffmpegBinary, this.AppendMovie.TsPath, this.Movie.TsPath, this.RawOption.ToString(), this.OutputFile)
	systemWithoutReturn(this, command)
	// fmt.Println(strings.Replace(system(command), "\n", "", -1))
	systemWithoutReturn(this, "rm -rf "+this.Movie.TsPath)
	systemWithoutReturn(this, "rm -rf "+this.AppendMovie.TsPath)
}

func (this *Transcoder) applyTranscoderOption() {
	this.TranscoderOption.Validate = true
	if this.Movie.calculatedAspectRatio() == 0 {
		return
	}
	if this.TranscoderOption.PreserveAspectRatio == "width" {
		newHeight := math.Ceil(float64(this.RawOption.Width()) / this.Movie.calculatedAspectRatio())
		this.RawOption.Resolution = fmt.Sprintf("%vx%v", this.RawOption.Width(), newHeight)
	} else if this.TranscoderOption.PreserveAspectRatio == "height" {
		newWidth := math.Ceil(float64(this.RawOption.Height()) / this.Movie.calculatedAspectRatio())
		this.RawOption.Resolution = fmt.Sprintf("%vx%v", newWidth, this.RawOption.Height())
	}
}

func (this *Transcoder) transcode(inputFile *string, outputFile *string) {
	ffmpegBinary := FfmpegBinary()
	command := fmt.Sprintf("%s -y -i %s %s %s ", ffmpegBinary, *inputFile, this.RawOption.ToString(), *outputFile)
	systemWithoutReturn(this, command)
	// strings.Replace(system(command), "\n", "", -1)
}

func (this *Transcoder) transcodeMovie() {
	ffmpegBinary := FfmpegBinary()
	command := fmt.Sprintf("%s -y -i %s %s %s", ffmpegBinary, this.Movie.Path, this.RawOption.ToString(), this.OutputFile)
	systemWithoutReturn(this, command)
	// strings.Replace(system(command), "\n", "", -1)
}
