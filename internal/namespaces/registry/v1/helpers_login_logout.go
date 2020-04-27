package registry

import "github.com/scaleway/scaleway-sdk-go/scw"

type program string
type programs []program

const (
	docker = program("docker")
	podman = program("podman")
)

var (
	availablePrograms = programs{docker, podman}
)

func (p programs) StringArray() []string {
	var res []string
	for _, prog := range p {
		res = append(res, string(prog))
	}
	return res
}

func getRegistryEndpoint(region scw.Region) string {
	return "rg." + region.String() + ".scw.cloud"
}
