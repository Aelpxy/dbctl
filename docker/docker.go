package docker

import (
	"bufio"
	"fmt"
	"os/exec"
)

func Delete_Container(id string) {
	args := []string{"rm", id, "-f"}
	Command_Exec("docker", args)
}

func Command_Exec(bin string, arg []string) {
	cmd := exec.Command(bin, arg...)
	r, _ := cmd.StdoutPipe()

	cmd.Stderr = cmd.Stdout

	done := make(chan struct{})
	scanner := bufio.NewScanner(r)

	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
		}

		done <- struct{}{}
	}()

	cmd.Start()

	<-done

	cmd.Wait()
}
