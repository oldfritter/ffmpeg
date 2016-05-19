package ffmpeg

import (
	"strings"
)

var ffmpegBinary, ffprobeBinary string

func FfmpegBinary() string {
	if ffmpegBinary == "" {
		ffmpegBinary = strings.Replace(system("which ffmpeg"), "\n", "", -1)
	}
	return ffmpegBinary
}

func FfprobeBinary() string {
	if ffprobeBinary == "" {
		ffprobeBinary = strings.Replace(system("which ffprobe"), "\n", "", -1)
	}
	return ffprobeBinary
}
