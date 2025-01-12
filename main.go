package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func getAgents() ([]string, error) {
	ctx := context.Background()
	client, clientErr := client.NewClientWithOpts(client.FromEnv)
	if clientErr != nil {
		return nil, clientErr
	}
	defer client.Close()

	// tasks, tasksErr := client.TaskList(ctx, types.TaskListOptions{
	// 	Filters: filters.NewArgs(
	// 		filters.KeyValuePair{Key: "desired-state", Value: "running"},
	// 		filters.KeyValuePair{Key: "label", Value: "syncthing.agent=true"},
	// 	),
	// })
	// if tasksErr != nil {
	// 	return nil, tasksErr
	// }

	tasks, tasksErr := client.ContainerList(ctx, container.ListOptions{
		Filters: filters.NewArgs(
		// filters.KeyValuePair{Key: "desired-state", Value: "running"},
		// filters.KeyValuePair{Key: "label", Value: "syncthing.agent=true"},
		),
	})
	if tasksErr != nil {
		return nil, tasksErr
	}

	result := []string{}
	for _, task := range tasks {
		// result = append(result, task.NodeID)
		fmt.Println(task.NetworkSettings.Networks["bridge"].IPAddress)
		result = append(result, task.HostConfig.NetworkMode)
	}
	return result, nil
}

func getDevices() ([]string, error) {
	resp, respErr := http.Get("")
	if respErr != nil {
		return nil, respErr
	}
	defer resp.Body.Close()

	return nil, nil
}

func main() {
	agents, agentsErr := getAgents()
	if agentsErr != nil {
		log.Fatal(agentsErr)
	}

	devices, devicesErr := getDevices()
	if devicesErr != nil {
		log.Fatal(devicesErr)
	}

	fmt.Println(nodes)
	fmt.Println(devices)
}
