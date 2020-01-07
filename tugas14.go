package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	if runtime.GOOS == "windows" {
		var output1, _ = exec.Command(`IPCONFIG`).Output()
		fmt.Println("-> IPCONFIG", string(output1))
		var output3, _ = exec.Command(`SYSTEMINFO`).Output()
		fmt.Println("-> SYSTEMINFO", string(output3))
		var output2, _ = exec.Command(`HELP`).Output()
		fmt.Println("-> HELP", string(output2))
	} else {
		fmt.Println("Error, Ini Bukan OS Windows")
	}
}
