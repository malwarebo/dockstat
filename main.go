package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/fatih/color"
)

const nameColumnWidth = 80

func main() {
    // Initialize Docker client
    cli, err := client.NewClientWithOpts(client.FromEnv)
    if err != nil {
        log.Fatal(err)
    }
    ctx := context.Background()

    // Parse command-line arguments
    if len(os.Args) == 1 {
        listContainers(cli, ctx)
    } else if len(os.Args) == 3 {
        switch os.Args[1] {
        case "kill":
            killContainer(cli, ctx, os.Args[2])
        case "log":
            showLogs(cli, ctx, os.Args[2])
        default:
            printUsage()
        }
    } else {
        printUsage()
    }
}

func listContainers(cli *client.Client, ctx context.Context) {
    containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
    if err != nil {
        log.Fatal(err)
    }

    if len(containers) == 0 {
        fmt.Println("No running containers found")
        return
    }

    fmt.Println(color.YellowString("| %-13s | %-80s | %-20s |", "ID", "Name", "Status"))
    fmt.Println(strings.Repeat("-", 118))

    for _, container := range containers {
        containerID := container.ID[:10]
        containerName := format(strings.TrimPrefix(container.Names[0], "/"), nameColumnWidth)
        containerStatus := container.Status

        var statusColor *color.Color
        if container.State == "running" {
            statusColor = color.New(color.FgGreen)
        } else {
            statusColor = color.New(color.FgRed)
        }

        fmt.Printf("| %-13s | %-80s | ", containerID, containerName)
        statusColor.Printf("%-20s", containerStatus)
        fmt.Println(" |")
    }
}

func killContainer(cli *client.Client, ctx context.Context, containerID string) {
    err := cli.ContainerKill(ctx, containerID, "SIGKILL")
    if err != nil {
        log.Fatalf("Error killing container %s: %v", containerID, err)
    }
    fmt.Printf("Container %s killed\n", containerID)
}

func showLogs(cli *client.Client, ctx context.Context, containerID string) {
    // Use types.ContainerLogsOptions explicitly
    options := types.ContainerLogsOptions{
        ShowStdout: true,
        ShowStderr: true,
        Follow:     false,
    }
    logs, err := cli.ContainerLogs(ctx, containerID, options)
    if err != nil {
        log.Fatalf("Error getting logs for container %s: %v", containerID, err)
    }
    defer logs.Close()

    _, err = stdcopy.StdCopy(os.Stdout, os.Stderr, logs)
    if err != nil {
        log.Fatalf("Error reading logs: %v", err)
    }
}

func printUsage() {
    fmt.Println("Usage:")
    fmt.Println("  dockstat          : list running containers")
    fmt.Println("  dockstat kill <container_id> : kill a running container")
    fmt.Println("  dockstat log <container_id>  : show logs of a container")
}

func format(s string, length int) string {
    if len(s) > length {
        return s[:length]
    }
    return s + strings.Repeat(" ", length-len(s))
}