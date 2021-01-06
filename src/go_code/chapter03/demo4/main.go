package main

import (
    "os/exec"
    "syscall"
)

func main() {
    // 有GUI调用
    exec.Command(`cmd`, `/c`, `start`, `https://www.jianshu.com`).Start()

    // 无GUI调用
    cmd := exec.Command(`cmd`, `/c`, `start`, `https://www.jianshu.com`)
    cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
    cmd.Start()
    
}
