package ffmpeg

import (
	"fmt"
	"testing"
)

func TestEncodingOption(t *testing.T) {
	eo := EncodingOption{VideoCodec: "h264", Aspect: "copy", AudioCodec: "AAC"}
	fmt.Println(eo.ToString())
}
