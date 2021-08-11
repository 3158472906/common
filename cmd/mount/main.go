package main

import (
	"log"
	"syscall"
)


//sysmboll
func main() {
	err := syscall.Mount("test",
		"/opt/isecnet",
		"bind",syscall.MS_BIND,"")
	if err != nil {
		log.Fatal("mount err:",err)
	}

	err = syscall.Unmount("/opt/isecnet",syscall.MNT_DETACH)
	if err != nil {
		log.Fatal("umount err:",err)
	}

}
