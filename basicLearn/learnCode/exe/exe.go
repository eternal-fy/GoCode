package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// 1) os.StartProcess //
	/*********************/
	/* Linux: */
	env := os.Environ()
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	// 1st example: list files
	pid, err := os.StartProcess("C:\\Program Files (x86)\\Tencent\\QQ\\Bin\\QQScLauncher.exe", []string{"QQScLauncher"}, procAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err) //
		os.Exit(1)
	}
	fmt.Println("The process id is %v", pid)
	cmd := exec.Command("dir")
	fmt.Printf(" %v", cmd)

}
