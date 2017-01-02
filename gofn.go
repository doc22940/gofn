package gofn

import (
	docker "github.com/fsouza/go-dockerclient"
	"github.com/nuveo/gofn/provision"
)

// Run runs the designed image
func Run(contextDir, dockerFile, imageName string) (stdout string, err error) {
	client := provision.FnClient("")

	img, err := provision.FnFindImage(client, imageName)
	if err != nil {
		return
	}

	var image string
	var container *docker.Container

	if img.ID == "" {
		image, _ = provision.FnImageBuild(client, contextDir, dockerFile, imageName)
	} else {
		image = "gofn/" + (imageName)
	}

	container, err = provision.FnContainer(client, image)
	if err != nil {
		return
	}

	stdout = provision.FnRun(client, container.ID).String()

	provision.FnRemove(client, container.ID)
	if err != nil {
		return
	}
	return
}
