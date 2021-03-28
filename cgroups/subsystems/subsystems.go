package subsystems

type ResourceConfig struct {
	MemoryLimit string
	CpuShare    string
	CpuSet      string
}

// path is the position of a cgroup in hieracy which is in tree structure
type Subsystem interface {
	Name() string                               // name of the subsystem
	Set(path string, res *ResourceConfig) error // set the resource limit of a cgroup
	Apply(path string, pid int) error           // add a process into cgroup
	Remove(path string) error                   // remove cgroup
}

var (
	SubsystemsIns = []Subsystem{
		&CpusetSubSystem{},
		&MemorySubSystem{},
		&CpuSubSystem{},
	}
)
