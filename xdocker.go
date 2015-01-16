package main

import (
	"fmt"
	"github.com/samalba/dockerclient"
	"github.com/spf13/cobra"
	"os"
	"treeptik.fr/ascii"
	"treeptik.fr/commands/xkill"
)

func main() {

	var Force bool
	var xDockerHost = os.Getenv("XDOCKER_HOST")
	docker, _ := dockerclient.NewDockerClient(xDockerHost, nil)

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of xdocker",
		Long:  `All software has versions. This is xdocker's`,
		Run: func(cmd *cobra.Command, args []string) {
			ascii.DisplayMainLogo()
			fmt.Println("Extended docker client // Provided by Treeptik, the Cloud and Java Company")
		},
	}

	var rootCmd = &cobra.Command{Use: "xdocker"}
	xkill.InitCommands(rootCmd, docker)
	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().BoolVarP(&Force, "force", "v", false, "force action")

	rootCmd.Execute()
}
