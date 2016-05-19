package ffmpeg

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestMovie(t *testing.T) {
	path, _ := filepath.Abs("movies/awesome_movie.mov")
	movie := Movie{}
	err := movie.Initialize(&path)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(movie)
	}
}
