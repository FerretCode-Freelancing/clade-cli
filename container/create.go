package container

import (
	"log"

	"github.com/ferretcode-freelancing/clade-cli/request"
)

func Create(args []string, insecure bool) error {
	if len(args) != 3 {
		createHelp()
	}

	name := args[0]
	imageURL := args[1]
	image := args[2]

	containerRequest := ContainerRequest{
		Name:     name,
		ImageURL: imageURL,
		Image:    image,
	}

	response, err := request.MakeRequest("/containers/create", containerRequest, insecure)
	if err != nil {
		return err
	}

	log.Println(response)

	return nil
}

func createHelp() {
	log.Fatalf("\ncreate\n\t1. name: the name for the container\n\t2. image URL: the URL to the container image in a registry\n\t3. image: the name of the image to create the container from\n")
}
