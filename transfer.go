package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
)

//配置文件, 默认使用 images.cfg 即可
const CONF_FILE = "images.cfg"

//存储 Image
func saveImage(args []string, wg *sync.WaitGroup) {

	err := exec.Command("docker", args...).Run()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		panic(err)
	}

	wg.Done()
}

func main() {

	f, err := os.Open(CONF_FILE)

	defer f.Close()

	if err == nil {

		buff := bufio.NewReader(f)

		for {
			wg := new(sync.WaitGroup)

			line, err := buff.ReadString('\n')

			if err != nil || io.EOF == err {
				break
			}

			wg.Add(1)

			var imageName = strings.TrimSpace(line)
			var tarName = fmt.Sprintf("%s.tar", imageName)

			args := []string{
				"save",
				"-o",
				fmt.Sprintf("copy/%s", strings.Replace(tarName, "/", "_", -1)),
				imageName,
			}

			fmt.Printf("docker %s \n", strings.Join(args, " "))

			go saveImage(args, wg)
			wg.Wait()
		}
	}
}
