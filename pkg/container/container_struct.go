// Go-podman/pkg/container/container_struct.go
package container

import "time"

type Port struct {
	ContainerPort int    `json:"container_port"`
	HostIP        string `json:"host_ip"`
	HostPort      int    `json:"host_port"`
	Protocol      string `json:"protocol"`
	Range         int    `json:"range"`
}

type Size struct {
	RootFsSize int `json:"rootFsSize"`
	RwSize     int `json:"rwSize"`
}

type Labels struct {
	Property1 string `json:"property1"`
	Property2 string `json:"property2"`
}

type Namespaces struct {
	Cgroup string `json:"Cgroup"`
	Ipc    string `json:"Ipc"`
	Mnt    string `json:"Mnt"`
	Net    string `json:"Net"`
	Pidns  string `json:"Pidns"`
	User   string `json:"User"`
	Uts    string `json:"Uts"`
}

type Container struct {
	AutoRemove   bool        `json:"AutoRemove"`
	CIDFile      string      `json:"CIDFile"`
	Command      []string    `json:"Command"`
	Created      time.Time   `json:"Created"`
	CreatedAt    string      `json:"CreatedAt"`
	ExitCode     int         `json:"ExitCode"`
	Exited       bool        `json:"Exited"`
	ExitedAt     int         `json:"ExitedAt"`
	ExposedPorts interface{} `json:"ExposedPorts"` // Assuming ExposedPorts can be null or any type
	Id           string      `json:"Id"`
	Image        string      `json:"Image"`
	ImageID      string      `json:"ImageID"`
	IsInfra      bool        `json:"IsInfra"`
	Labels       Labels      `json:"Labels"`
	Mounts       []string    `json:"Mounts"`
	Names        []string    `json:"Names"`
	Namespaces   Namespaces  `json:"Namespaces"`
	Networks     []string    `json:"Networks"`
	Pid          int         `json:"Pid"`
	Pod          string      `json:"Pod"`
	PodName      string      `json:"PodName"`
	Ports        []Port      `json:"Ports"`
	Restarts     int         `json:"Restarts"`
	Size         Size        `json:"Size"`
	StartedAt    int         `json:"StartedAt"`
	State        string      `json:"State"`
	Status       string      `json:"Status"`
}
