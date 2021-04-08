package main

import (
	"My-Docker/cgroups"
	"My-Docker/cgroups/subsystems"
	"My-Docker/container"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
)

func Run(tty bool, comArray []string, res *subsystems.ResourceConfig, volume string) {
	parent, writePipe := container.NewParentProcess(tty, volume)
	if parent == nil {
		log.Errorf("New parent process error")
		return
	}
	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	// use mydocker-cgroup as cgroup name
	// config cgroup
	cgroupManager := cgroups.NewCgroupManager("My-Docker-cgroup")
	defer cgroupManager.Destroy()
	cgroupManager.Set(res)
	cgroupManager.Apply(parent.Process.Pid)

	sendInitCommand(comArray, writePipe)
	if tty {
		parent.Wait()
		// delete workspace
		mntURL := "/root/mnt"
		rootURL := "/root"
		container.DeleteWorkSpace(rootURL, mntURL, volume)
	}

}

func sendInitCommand(comArray []string, writePipe *os.File) {
	command := strings.Join(comArray, " ")
	log.Infof("command all is %s", command)
	writePipe.WriteString(command)
	writePipe.Close()
}
