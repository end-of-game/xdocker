package xremove

import (
	"fmt"
	"github.com/samalba/dockerclient"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

// global variable for this package
var docker *dockerclient.DockerClient

// Function to define commands
func InitCommands(rootCmd *cobra.Command, pdocker *dockerclient.DockerClient) {

	docker = pdocker

	var RemoveCmd = &cobra.Command{
		Use:   "remove",
		Short: "Remove containers",
		Long:  "Remove containers.",
		Run:   Remove,
	}

	var RemoveContainersCmd = &cobra.Command{
		Use:   "containers",
		Short: "Remove all containers",
		Long:  "Remove all containers with optional selection",
		Run:   RemoveContainers,
	}

	RemoveCmd.AddCommand(RemoveContainersCmd)
	rootCmd.AddCommand(RemoveCmd)
}

func Remove(cmd *cobra.Command, args []string) {
	fmt.Println("Command options are \n")
	fmt.Println("\txdocker remove containers")
	fmt.Println("\txdocker remove images")
	fmt.Println("")
}

func RemoveContainers(cmd *cobra.Command, args []string) {
	containers, err := docker.ListContainers(false, true, "")
	if err != nil {
		log.Fatal(err)
	}
	if len(containers) == 0 {
		fmt.Println("*************************")
		fmt.Println("*** Nothing to remove ***")
		fmt.Println("*************************")
		return
	}
	fmt.Println("")

	for _, c := range containers {
		containerInfo, _ := docker.InspectContainer(c.Id)
		currentStatus := ""
		if containerInfo.State.Running {
			currentStatus = "Running"
		} else if containerInfo.State.Paused {
			currentStatus = "Paused"
		} else if containerInfo.State.Ghost {
			currentStatus = "Ghost"
		} else {
			currentStatus = "Stopped"
		}
		fmt.Printf("[Ready to remove -- %s] %s \n", currentStatus, c.Names[0][1:])
	}

	fmt.Print("\n[Warning] Are you sure to remove them [Yes/No] : ")
	var input string
	fmt.Scanf("%s", &input)
	// if the answer is yes, we can kill the container
	if strings.EqualFold(input, "yes") || strings.EqualFold(input, "y") {
		fmt.Println("")
		var counter int
		for _, c := range containers {
			err = docker.RemoveContainer(c.Id, true)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Printf("[Removing] %s \n", c.Names[0][1:])
				counter++
			}
			fmt.Printf("[Removing : %s] %s \n", c.Names[0][1:])
		}
		fmt.Println("")
		fmt.Println(counter, "containers are removed")
		fmt.Println("")
	} else {
		fmt.Println("\nThanks to take time to think about...")
	}
}
