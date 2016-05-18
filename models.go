package ffmpeg

import (
	"time"
)

type Movie struct {
	Path         string
	TsPath       string
	Duration     float64
	Time         float64
	Bitrate      int64
	CreationTime time.Time
	Container    string
	Invalid      bool

	VideoStream  string
	VideoCodec   string
	VideoBitrate int64
	Colorspace   string
	Width        int
	Height       int
	Sar          string
	Dar          string
	FrameRate    string
	Rotation     int64

	AudioStream        string
	AudioCodec         string
	AudioBitrate       string
	AudioSampleRate    string
	AudioChannels      int
	AudioChannelLayout string
}

type Result struct {
	Streams []Stream
	Format  Format
	Error   string
}

type Stream struct {
	Index            int
	Codec_Name       string
	Codec_Long_Name  string
	Profile          string
	Codec_Type       string
	Codec_Time_Base  string
	Codec_Tag_String string
	Codec_Tag        string
	R_Frame_Rate     string
	Avg_Frame_Rate   string
	Time_Base        string
	Start_Pts        int
	Start_Time       string
	Duration_Ts      int
	Duration         string
	Bit_Rate         string
	Max_Bit_Rate     string
	Nb_Frames        string
	Disposition      Disposition
	Tags             Tag
	// video
	Width                int
	Height               int
	Coded_Width          int
	Coded_Height         int
	Sample_Aspect_Ratio  string
	Display_Aspect_Ratio string
	Pix_Fmt              string
	Level                int
	Chroma_Location      string
	Refs                 int
	Quarter_Sample       string
	Divx_Packed          string
	// audio
	Sample_Fmt      string
	Sample_Rate     string
	Channels        int
	Channel_Layout  string
	Bits_Per_Sample int
}

type Disposition struct {
	Default          int
	Dub              int
	Original         int
	Comment          int
	Lyrics           int
	Karaoke          int
	Forced           int
	Hearing_Impaired int
	Visual_Impaired  int
	Clean_Effects    int
	Attached_Pic     int
}

type Tag struct {
	Rotate        string
	Creation_Time string
	Language      string
	Handler_Name  string
}

type Format struct {
	Filename         string
	Nb_Streams       int
	Nb_Programs      int
	Format_Name      string
	Format_Long_Name string
	Start_Time       string
	Duration         string
	Size             string
	Bit_Rate         string
	Probe_Score      int
	Tags             FormatTag
}

type FormatTag struct {
	Major_Brand       string
	Minor_Version     string
	Compatible_Brands string
	Creation_Time     string
	Encoder           string
}
