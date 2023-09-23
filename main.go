package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/fatih/color"
)

const nameColumnWidth = 80

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	if len(containers) == 0 {
		log.Fatalf("No running containers found")
	}

	fmt.Println(color.YellowString("| %-13s | %-80s | %-20s |", "ID", "Name", "Status"))
	fmt.Println(strings.Repeat("-", 118))

	for _, container := range containers {
		containerID := container.ID[:10]
		containerName := format(container.Names[0], nameColumnWidth)
		containerStatus := container.Status

		// Check if the container is running or stopped and apply the color accordingly
		var statusColor *color.Color
		if container.State == "running" {
			statusColor = color.New(color.FgGreen)
		} else {
			statusColor = color.New(color.FgRed)
		}

		// Print the colored output in a table-like format
		fmt.Printf("| %-13s | %-80s | ", containerID, containerName)
		statusColor.Printf("%-20s", containerStatus)
		fmt.Println(" |")
	}
}
