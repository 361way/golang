/*
通过系统指令，获取系统相关数据
并通过time模块和for循环，实现不停的抓取系统数据。
*/
package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"
)

func Shellout(command string) (string, string, error) {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	return stdout.String(), stderr.String(), err
}

func main() {

	tc := time.Tick(30 * time.Second)
	for {

		mem, _, _ := Shellout("free -m")
		cpu, _, _ := Shellout("top -b -n1 | grep -B1 '%Cpu'")
		uptime, _, _ := Shellout("uptime")
		disk, _, _ := Shellout("df -hP")
		inode, _, _ := Shellout("df -hPi")

		fmt.Println(time.Now().Local())
		fmt.Println(mem)
		fmt.Println(cpu)
		fmt.Println(uptime)
		fmt.Println(disk)
		fmt.Println(inode)
		fmt.Println("################################################################")
		<-tc
	}

}

/*
也可以用如下格式运行
`
echo "## date"
date +'%Y-%m-%d %T'

echo "## meminfo"
free -m

echo "## cpuinfo"
top -b -n1 | grep -B1 '%Cpu'

echo "## uptime-load"
uptime

echo "## disk"
df -hP

echo "## inode"
df -hPi
`
sysinfo, _, _ := Shellout(cmd)
fmt.Println(sysinfo)
*/

