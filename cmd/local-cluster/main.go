package main

import (
	"fmt"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/olekukonko/tablewriter"
	"github.com/simonmckeon/local-cluster/internal/docker"
)

func main() {
	var containers []types.ContainerJSON
	var containerDetails []docker.ContainerDetail

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		containerID, err := docker.CreateContainer(cli, "centos:7")
		if err != nil {
			panic(err)
		}
		fmt.Printf("Container created: %s\n", containerID)

		container, err := docker.StartContainer(cli, containerID)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Container started: %s\n", containerID)

		containers = append(containers, container)
	}

	for _, container := range containers {
		containerDetails = append(containerDetails, docker.GetContainerDetails(cli, &container))
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Image", "IP address"})
	for _, containerDetail := range containerDetails {
		table.Append([]string{
			containerDetail.ID,
			containerDetail.Name,
			containerDetail.Image,
			containerDetail.IPAddress,
		})
	}
	fmt.Println("CLUSTER DETAILS ---------------------")
	table.Render()
}
