package container

import (
	"os"
	"syscall"

	"github.com/Sirupsen/logrus"
)

func RunContainerInitProcess(command string, args []string) error {
	logrus.Infof("command %s", command)

	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	// I use proc proc /proc here while textbook uses proc /proc proc
	// Using textbook one will make /proc unmounted when exit from container
	syscall.Mount("proc", "proc", "/proc", uintptr(defaultMountFlags), "")
	argv := []string{command}
	if err := syscall.Exec(command, argv, os.Environ()); err != nil {
		logrus.Errorf(err.Error())
	}
	return nil
}
