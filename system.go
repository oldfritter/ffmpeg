package ffmpeg

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func system(command string) string {
	fmt.Println("Running command...\n", command)
	cmd := exec.Command("/bin/sh", "-c", command)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	_, err = cmd.Process.Wait()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}
	return out.String()
}
