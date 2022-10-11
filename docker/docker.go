package docker

import (
	"bufio"
	"log"
	"os"
	"os/exec"
)

func Pull_Image(imageName string) {
	cmd := exec.Command("docker", "pull", imageName)
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Fatal(err)
	}

	cmd.Start()

	buf := bufio.NewReader(stdout)
	for {
		line, _, _ := buf.ReadLine()
		if line == nil {
			log.Println("Successfully pulled database image")
			os.Exit(0)
		}
		log.Println(string(line))
	}
}