package main

import (
	"github.com/dsmontoya/axolog/client"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	rootCmd := &cobra.Command{
		Use:   "axolog",
		Short: "Docker log router",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			client, err := client.New(args[0])
			if err != nil {
				panic(err)
			}
			client.RegisterContainers()

			if err != nil {
				panic(err)
			}
			if err := client.ReadLogs(); err != nil {
				panic(err)
			}

			if err := client.ListenEvents(); err != nil {
				panic(err)
			}
		},
	}

	rootCmd.Execute()
}
