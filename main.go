package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/client"
	// "github.com/docker/docker/api/types/container"
	// "github.com/docker/docker/api/types/events"
	// "github.com/docker/docker/client"
)

func getIPs() ([]string, error) {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	defer apiClient.Close()

	return nil, nil
}

func main() {

	ips, _ := getIPs(apiClient)
	ids, _ := getIDs()

	listContainers(apiClient)

	// services, err := apiClient.ServiceList(context.Background(), types.ServiceListOptions{})
	// if err != nil {
	// 	panic(err)
	// }

	containers, err := apiClient.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		panic(err)
	}

	for _, ctr := range containers {
		fmt.Printf("%s %s (status: %s)\n", ctr.ID, ctr.Image, ctr.Status)
	}

	events, errs := apiClient.Events(context.Background(), events.ListOptions{})
	for {
		select {
		case err := <-errs:
			fmt.Println(err)
		case event := <-events:
			fmt.Println("------------------")
			fmt.Println(event.Type)
			fmt.Println(event.Action)
			fmt.Println(event.Actor)
			fmt.Println("------------------")
		}
	}
}
