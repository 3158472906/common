package main

import (
	"log"
	"syscall"
)


func main() {
	err := syscall.Mount("test/test.file",
		"/opt/isecnet/test.file",
		"bind",syscall.MS_BIND,"")
	if err != nil {
		log.Fatal("mount err:",err)
	}

	//err = syscall.Unmount("/opt/isecnet/test.file",syscall.MNT_DETACH)
	//if err != nil {
	//	log.Fatal("umount err:",err)
	//}

}
