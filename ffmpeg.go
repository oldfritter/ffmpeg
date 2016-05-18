package ffmpeg

import (
	"strings"
)

func FfmpegBinary() string {
	return strings.Replace(system("which ffmpeg"), "\n", "", -1)
}

func FfprobeBinary() string {
	return strings.Replace(system("which ffprobe"), "\n", "", -1)
}
