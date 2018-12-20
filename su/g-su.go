package main

import (
	"fmt"

	gexpect "github.com/ThomasRooney/gexpect"
)

func main() {
	child, err := gexpect.Spawn("su - zabbix")
	if err != nil {
		panic(err)
	}

	child.Expect("Password")
	child.SendLine("zabbixpasswd")
	child.Interact()

}

//缺点较多，建议使用syscall通过底层uid切换，并配合s权限位实现。
