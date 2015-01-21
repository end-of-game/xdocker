package xkill

import (
	"fmt"
	"github.com/samalba/dockerclient"
	"github.com/spf13/cobra"
	"log"
	"regexp"
	"strings"
)

// global variable for this package
var docker *dockerclient.DockerClient

// Function to define commands
func InitCommands(rootCmd *cobra.Command, pdocker *dockerclient.DockerClient) {

	docker = pdocker

	var allCmd = &cobra.Command{
		Use:   "all",
		Short: "Kill all containers",
		Long:  "Kill all containers",
		Run:   KillAll,
	}

	var regexpCmd = &cobra.Command{
		Use:   "rex",
		Short: "Kill containers according to regular expression",
		Long:  "Kill containers according to regular expression",
		Run:   KillWithRegex,
	}

	var cmdKill = &cobra.Command{
		Use:   "kill [all|rex]",
		Short: "Kill containers",
		Long: `Kill container.
        The command accepts 'all' and 'rex' as subcommand.
        `,
		Run: Kill,
	}

	cmdKill.AddCommand(allCmd)
	cmdKill.AddCommand(regexpCmd)
	rootCmd.AddCommand(cmdKill)
}

func Kill(cmd *cobra.Command, args []string) {
	fmt.Println("Command options are \n")
	fmt.Println("\txdocker kill all")
	fmt.Println("\txdocker kill rex [pattern]")
	fmt.Println("")
}

func KillAll(cmd *cobra.Command, args []string) {
	containers, err := docker.ListContainers(false, true, "")
	if err != nil {
		log.Fatal(err)
	}
	if len(containers) == 0 {
		fmt.Println("***********************")
		fmt.Println("*** Nothing to hunt ***")
		fmt.Println("***********************")
		return
	}
	fmt.Println("")

	for _, c := range containers {
		fmt.Printf("[Selection] %s \n", c.Names[0][1:])
	}

	fmt.Print("\n[Warning] Are you sure to kill them [Yes/No] : ")
	var input string
	fmt.Scanf("%s", &input)
	// if the answer is yes, we can kill the container
	if strings.EqualFold(input, "yes") || strings.EqualFold(input, "y") {
		fmt.Println("")
		var counter int
		for _, c := range containers {
			err = docker.KillContainer(c.Id, "")
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Printf("[Shooting] %s \n", c.Names[0][1:])
				counter++
			}
		}
		fmt.Println("")
		fmt.Println(counter, "containers are shooted")
		fmt.Println("")
	} else {
		fmt.Println("\nThanks to take time to think about...")
	}
}

func KillWithRegex(cmd *cobra.Command, args []string) {
	containers, err := docker.ListContainers(false, true, "")
	if err != nil {
		log.Fatal(err)
	}
	if len(containers) == 0 {
		fmt.Println("***********************")
		fmt.Println("*** Nothing to hunt ***")
		fmt.Println("***********************")
		return
	}
	fmt.Println("")

	var counter int
	r, _ := regexp.Compile(args[0])
	for _, c := range containers {
		if r.MatchString(c.Names[0]) {
			err = docker.KillContainer(c.Id, "")
			if err != nil {
				log.Fatal(err)
			} else {
				counter++
				fmt.Printf("[Shooting] %s is killed\n", c.Names[0][1:])
			}
		}
	}
	if counter == 0 {
		fmt.Println("***********************************************")
		fmt.Println("*** Nothing to hunt for pattern : ", args[0])
		fmt.Println("***********************************************")
		return
	} else {

	}
	fmt.Println("")
	fmt.Println(counter, "containers are shooted")
	fmt.Println("")
}
