package ffmpeg

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func (this *Movie) Initialize(path *string) error {
	_, err := os.Stat(*path)
	if os.IsNotExist(err) {
		return fmt.Errorf("the file %s does not exist", *path)
	}
	this.Path = *path
	suffix := regexp.MustCompile("[avi|mp4|flv|mkv|mov]+$")
	this.TsPath = suffix.ReplaceAllString(*path, "ts")
	ffprobe_binary := FfprobeBinary()
	command := fmt.Sprintf("%s -i %s -print_format json -show_format -show_streams -show_error", ffprobe_binary, *path)
	std_output := system(command)
	var metadata Result
	err = json.Unmarshal([]byte(std_output), &metadata)
	var videoStreams []Stream
	var audioStreams []Stream
	if err != nil {
		this.Duration = 0
	} else {
		for _, stream := range metadata.Streams {
			if stream.Codec_Type == "video" {
				videoStreams = append(videoStreams, stream)
			} else if stream.Codec_Type == "audio" {
				audioStreams = append(audioStreams, stream)
			}
		}
	}
	this.Container = metadata.Format.Format_Name
	this.Duration, _ = strconv.ParseFloat(metadata.Format.Duration, 64)
	this.Time, _ = strconv.ParseFloat(metadata.Format.Start_Time, 64)
	this.CreationTime, _ = time.Parse("2006-01-03 15:04:05", metadata.Format.Tags.Creation_Time)
	this.Bitrate, _ = strconv.ParseInt(metadata.Format.Bit_Rate, 10, 64)
	if len(videoStreams) > 0 {
		videoStream := videoStreams[0]
		this.VideoCodec = videoStream.Codec_Name
		this.Colorspace = videoStream.Pix_Fmt
		this.Width = videoStream.Width
		this.Height = videoStream.Height
		this.VideoBitrate, _ = strconv.ParseInt(videoStream.Bit_Rate, 10, 64)
		this.Sar = videoStream.Sample_Aspect_Ratio
		this.Dar = videoStream.Display_Aspect_Ratio
		if videoStream.Avg_Frame_Rate != "0/0" {
			this.FrameRate = videoStream.Avg_Frame_Rate
		}
		this.VideoStream = fmt.Sprintf("%s (%s) (%s/%s), %s, %s [SAR %s DAR %s]", videoStream.Codec_Name, videoStream.Profile, videoStream.Codec_Tag_String, videoStream.Codec_Tag, this.Colorspace, this.Resolution(), this.Sar, this.Dar)
		if videoStream.Tags.Rotate != "" {
			this.Rotation, _ = strconv.ParseInt(videoStream.Tags.Rotate, 10, 64)
		}
	}
	if len(audioStreams) > 0 {
		audioStream := audioStreams[0]
		this.AudioChannels = audioStream.Channels
		this.AudioCodec = audioStream.Codec_Name
		this.AudioSampleRate = audioStream.Sample_Rate
		this.AudioBitrate = audioStream.Bit_Rate
		this.AudioChannelLayout = audioStream.Channel_Layout
		this.AudioStream = fmt.Sprintf("%s (%s/%s), %s Hz, %s, %s, %s bit/s", this.AudioCodec, audioStream.Codec_Tag_String, audioStream.Codec_Tag, this.AudioSampleRate, this.AudioChannelLayout, audioStream.Sample_Fmt, this.AudioBitrate)
	}
	if metadata.Error != "" {
		this.Invalid = true
	}
	return nil
}

func (this *Movie) Valid() bool {
	if this.Invalid {
		return false
	} else {
		return true
	}
}

func (this *Movie) TWidth() int {
	if this.Rotation == 0 || this.Rotation == 180 {
		return this.Width
	} else {
		return this.Height
	}
}

func (this *Movie) THeight() int {
	if this.Rotation == 0 || this.Rotation == 180 {
		return this.Height
	} else {
		return this.Width
	}
}

func (this *Movie) Resolution() string {
	return fmt.Sprintf("%dx%d", this.TWidth(), this.THeight())
}

func (this *Movie) CalculatedAspectRatio() float64 {
	re := this.aspectFromDar()
	if re != 0 {
		return re
	}
	return this.aspectFromDimensions()
}

func (this *Movie) CalculatedPixelAspectRatio() float64 {
	if re := this.aspectFromSar(); re != 0 {
		return re
	}
	return 1
}

func (this *Movie) Size() int64 {
	fileInfo, _ := os.Stat(this.Path)
	return fileInfo.Size()
}

func (this *Movie) Transcode(outputFile *string, option *EncodingOption, transcoderOption *TranscoderOption) {
	transcoder := Transcoder{}
	transcoder.Initialize(this, outputFile, option, transcoderOption)
	transcoder.Run()
}

func (this *Movie) Append(outputFile *string, option *EncodingOption, transcoderOption *TranscoderOption) {
	transcoder := Transcoder{}
	transcoder.Initialize(this, outputFile, option, transcoderOption)
	transcoder.Append()
}

func (this *Movie) aspectFromSar() float64 {
	if this.Sar == "" {
		return 0
	}
	dars := strings.Split(this.Dar, ":")
	w, _ := strconv.ParseFloat(dars[0], 64)
	h, _ := strconv.ParseFloat(dars[1], 64)

	if this.Rotation == 0 || this.Rotation == 180 {
		return w / h
	} else {
		return h / w
	}
	return 0
}

func (this *Movie) aspectFromDar() float64 {
	if this.Dar == "" {
		return 0
	}
	dars := strings.Split(this.Dar, ":")
	w, _ := strconv.ParseFloat(dars[0], 64)
	h, _ := strconv.ParseFloat(dars[1], 64)

	if this.Rotation == 0 || this.Rotation == 180 {
		return w / h
	} else {
		return h / w
	}
	return 0
}

func (this *Movie) aspectFromDimensions() float64 {
	re := this.TWidth() / this.THeight()
	return float64(re)
}

func (this *Movie) calculatedAspectRatio() float64 {
	re := this.aspectFromDar()
	if re == 0 {
		afd := this.aspectFromDimensions()
		if afd == 0 {
			return 0
		}
		return afd
	}
	return re
}
