package main

import (
	"fmt"
	"github.com/samalba/dockerclient"
	"github.com/spf13/cobra"
	"log"
	"regexp"
	"treeptik.fr/ascii"
)

func main() {

	ascii.DisplayMainLogo()

	var Force bool
	docker, _ := dockerclient.NewDockerClient("tcp://192.168.50.4:4243", nil)
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of xdocker",
		Long:  `All software has versions. This is xdocker's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Extended docker client // Provided by Treeptik, the Cloud and Java Company")
		},
	}

	var allCmd = &cobra.Command{
		Use:   "all",
		Short: "Kill all containers",
		Long:  "Kill all containers with confirmation",
		Run: func(cmd *cobra.Command, args []string) {
			//fmt.Println("Kill them all !", Force)
			// Get only running containers
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
		},
	}

	var regexpCmd = &cobra.Command{
		Use:   "rex",
		Short: "Kill containers through regular expression",
		Long:  "Kill containers through regular expression",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Kill them with a regular expression", Force)
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
		},
	}

	var cmdKill = &cobra.Command{
		Use:   "kill [all|rex]",
		Short: "Kill containers",
		Long: `Kill container.
        The command accepts 'all' and 'rex' as subcommand.
        `,
		Run: func(cmd *cobra.Command, args []string) {
			// Get only running containers
			containers, err := docker.ListContainers(false, true, "")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("")
			if len(containers) == 0 {
				fmt.Println("***********************")
				fmt.Println("*** Nothing to hunt ***")
				fmt.Println("***********************\n")
				return
			}
			for _, c := range containers {
				fmt.Printf("[Simulation] kill %s\n", c.Names[0][1:])
			}
			fmt.Println("Command options are \n")
			fmt.Println("\txdocker kill all")
			fmt.Println("\txdocker kill rex [pattern]")
			fmt.Println("")
		},
	}

	var rootCmd = &cobra.Command{Use: "xdocker"}
	rootCmd.AddCommand(cmdKill, versionCmd)
	cmdKill.AddCommand(allCmd)
	cmdKill.AddCommand(regexpCmd)

	rootCmd.PersistentFlags().BoolVarP(&Force, "force", "v", false, "force action")

	rootCmd.Execute()
}
