package util

import (
	"fmt"
	"os/exec"
	"runtime"
)

// OpenExplore 打开浏览器
func OpenExplore(urt string) {
	var cmd string
	var args []string
	switch runtime.GOOS {
	case "windows":
		cmd = "explorer"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default:
		cmd = "xdg-open"
	}
	args = append(args, urt)

	err := exec.Command(cmd, args...).Start()
	if err != nil {
		fmt.Println("浏览器打开失败:", err)
	} else {
		fmt.Println("浏览器打开成功")
	}
}
