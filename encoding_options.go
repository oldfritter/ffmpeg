package ffmpeg

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type EncodingOption struct {
	AppendMovie           string
	Aspect                string
	VideoCodec            string
	FrameRate             string
	Resolution            string
	VideoBitrate          string
	AudioCodec            string
	AudioBitrate          string
	AudioSampleRate       string
	AudioChannels         string
	VideoMaxBitrate       string
	VideoMinBitrate       string
	BufferSize            string
	VideoBitrateTolerance string
	Threads               string
	Target                string
	Duration              string
	VideoPreset           string
	AudioPreset           string
	FilePreset            string
	KeyframeInterval      string
	SeekTime              string
	Vframes               string
	Quality               string
	X264Vprofile          string
	X264VPreset           string
	Watermark             string
}

func (this *EncodingOption) ToString() string {
	var optionString string
	funcs := []string{"Aspect", "VideoCodec", "FrameRate", "Resolution", "VideoBitrate", "AudioCodec", "AudioBitrate", "AudioSampleRate", "AudioChannels", "VideoMaxBitrate", "VideoMinBitrate", "BufferSize", "VideoBitrateTolerance", "Threads", "Target", "Duration", "VideoPreset", "AudioPreset", "FilePreset", "KeyframeInterval", "SeekTime", "Vframes", "Quality", "X264Vprofile", "X264VPreset", "Watermark"}
	for _, name := range funcs {
		re := reflect.ValueOf(this).MethodByName(fmt.Sprintf("Convert%s", name)).Call([]reflect.Value{})
		if re[0].String() != "" {
			optionString += " "
			optionString += re[0].String()
		}
	}
	return optionString
}

func (this *EncodingOption) ConvertAspect() string {
	if this.canCalculateAspect() {
		res := strings.Split(this.Resolution, "x")
		if len(res) > 1 {
			width, _ := strconv.ParseFloat(res[0], 64)
			height, _ := strconv.ParseFloat(res[1], 64)
			return fmt.Sprintf("-aspect %f", width/height)
		}
	}
	if this.Aspect != "" {
		return fmt.Sprintf("-aspect %s", this.Aspect)
	}
	return ""
}

func (this *EncodingOption) ConvertVideoCodec() string {
	if this.VideoCodec != "" {
		return fmt.Sprintf("-vcodec %s", this.VideoCodec)
	}
	return ""
}
func (this *EncodingOption) ConvertFrameRate() string {
	if this.FrameRate != "" {
		return fmt.Sprintf("-r %s", this.FrameRate)
	}
	return ""
}

func (this *EncodingOption) ConvertResolution() string {
	if this.Resolution != "" {
		return fmt.Sprintf("-s %s", this.Resolution)
	}
	return ""
}

func (this *EncodingOption) ConvertVideoBitrate() string {
	if this.VideoBitrate != "" {
		return fmt.Sprintf("-b:v %s", this.VideoBitrate)
	}
	return ""
}

func (this *EncodingOption) ConvertAudioCodec() string {
	if this.AudioCodec != "" {
		return fmt.Sprintf("-acodec %s", this.AudioCodec)
	}
	return ""
}

func (this *EncodingOption) ConvertAudioBitrate() string {
	if this.AudioBitrate != "" {
		return fmt.Sprintf("-b:a %s", this.AudioBitrate)
	}
	return ""
}

func (this *EncodingOption) ConvertAudioSampleRate() string {
	if this.AudioSampleRate != "" {
		return fmt.Sprintf("-ar %s", this.AudioSampleRate)
	}
	return ""
}

func (this *EncodingOption) ConvertAudioChannels() string {
	if this.AudioChannels != "" {
		return fmt.Sprintf("-ac %s", this.AudioChannels)
	}
	return ""
}

func (this *EncodingOption) ConvertVideoMaxBitrate() string {
	if this.VideoMaxBitrate != "" {
		return fmt.Sprintf("-maxrate %s", this.kFormat(this.VideoMaxBitrate))
	}
	return ""
}

func (this *EncodingOption) ConvertVideoMinBitrate() string {
	if this.VideoMinBitrate != "" {
		return fmt.Sprintf("-minrate %s", this.kFormat(this.VideoMinBitrate))
	}
	return ""
}

func (this *EncodingOption) ConvertBufferSize() string {
	if this.BufferSize != "" {
		return fmt.Sprintf("-bufsize %s", this.kFormat(this.BufferSize))
	}
	return ""
}

func (this *EncodingOption) ConvertVideoBitrateTolerance() string {
	if this.VideoBitrateTolerance != "" {
		return fmt.Sprintf("-bt %s", this.kFormat(this.VideoBitrateTolerance))
	}
	return ""
}

func (this *EncodingOption) ConvertThreads() string {
	if this.Threads != "" {
		return fmt.Sprintf("-threads %s", this.Threads)
	}
	return ""
}

func (this *EncodingOption) ConvertTarget() string {
	if this.Target != "" {
		return fmt.Sprintf("-target %s", this.Target)
	}
	return ""
}

func (this *EncodingOption) ConvertDuration() string {
	if this.Duration != "" {
		return fmt.Sprintf("-t %s", this.Duration)
	}
	return ""
}

func (this *EncodingOption) ConvertVideoPreset() string {
	if this.VideoPreset != "" {
		return fmt.Sprintf("-vpre %s", this.VideoPreset)
	}
	return ""
}

func (this *EncodingOption) ConvertAudioPreset() string {
	if this.AudioPreset != "" {
		return fmt.Sprintf("-apre %s", this.AudioPreset)
	}
	return ""
}

func (this *EncodingOption) ConvertFilePreset() string {
	if this.FilePreset != "" {
		return fmt.Sprintf("-fpre %s", this.FilePreset)
	}
	return ""
}

func (this *EncodingOption) ConvertKeyframeInterval() string {
	if this.KeyframeInterval != "" {
		return fmt.Sprintf("-g %s", this.KeyframeInterval)
	}
	return ""
}

func (this *EncodingOption) ConvertSeekTime() string {
	if this.SeekTime != "" {
		return fmt.Sprintf("-ss %s", this.SeekTime)
	}
	return ""
}

func (this *EncodingOption) ConvertScreenshot(value string) string {
	if value == "" {
		return ""
	}
	if this.Vframes == "" {
		this.Vframes = "-vframes 1 "
	}
	return fmt.Sprintf(" %s-f image2 ", this.Vframes)
}

func (this *EncodingOption) ConvertQuality() string {
	if this.Quality != "" {
		return fmt.Sprintf("-q:v %s", this.Quality)
	}
	return ""
}

func (this *EncodingOption) ConvertVframes() string {
	if this.Vframes != "" {
		return fmt.Sprintf("-vframes %s", this.Vframes)
	}
	return ""
}

func (this *EncodingOption) ConvertX264Vprofile() string {
	if this.X264Vprofile != "" {
		return fmt.Sprintf("-vprofile %s", this.X264Vprofile)
	}
	return ""
}

func (this *EncodingOption) ConvertX264VPreset() string {
	if this.X264VPreset != "" {
		return fmt.Sprintf("-preset %s", this.X264VPreset)
	}
	return ""
}

func (this *EncodingOption) ConvertWatermark() string {
	if this.Watermark != "" {
		return fmt.Sprintf("-i %s", this.Watermark)
	}
	return ""
}

func (this *EncodingOption) ConvertWatermarkFilter(value map[string]string) string {
	if value["position"] == "LT" {
		return fmt.Sprintf("-filter_complex 'scale=%s,overlay=x=%s:y=%s'", this.Resolution, value["padding_x"], value["padding_y"])
	} else if value["position"] == "RT" {
		return fmt.Sprintf("-filter_complex 'scale=%s,overlay=x=main_w-overlay_w-%s:y=%s'", this.Resolution, value["padding_x"], value["padding_y"])
	} else if value["position"] == "LB" {
		return fmt.Sprintf("-filter_complex 'scale=%s,overlay=x=%s:y=main_h-overlay_h-%s'", this.Resolution, value["padding_x"], value["padding_y"])
	} else if value["position"] == "RB" {
		return fmt.Sprintf("-filter_complex 'scale=%s,overlay=x=main_w-overlay_w-%s:y=main_h-overlay_h-%s'", this.Resolution, value["padding_x"], value["padding_y"])
	}
	return ""
}

func (this *EncodingOption) ConvertCustom(value string) string {
	return value
}

func (this *EncodingOption) kFormat(value string) string {
	re, _ := regexp.MatchString("k", value)
	if re == true {
		return value
	}
	return fmt.Sprintf("%sk", value)
}

func (this *EncodingOption) canCalculateAspect() bool {
	if this.Aspect == "" || this.Resolution != "" {
		return true
	}
	return false
}

func (this *EncodingOption) Width() int64 {
	res := strings.Split(this.Resolution, "x")
	width, _ := strconv.ParseInt(res[0], 10, 64)
	return width
}

func (this *EncodingOption) Height() int64 {
	res := strings.Split(this.Resolution, "x")
	height, _ := strconv.ParseInt(res[1], 10, 64)
	return height
}
