package ffmpeg

import (
	"bytes"
	"log"
	"os/exec"
)

func system(command string) string {
	cmd := exec.Command("/bin/sh", "-c", command)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return out.String()
}
