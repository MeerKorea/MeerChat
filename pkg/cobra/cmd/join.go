/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/jjwow73/MeerChat/pkg/params"
	"github.com/spf13/cobra"
)

// joinCmd represents the join command
var joinCmd = &cobra.Command{
	Use:   "join",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ip := getFlagString(cmd, "ip")
		port := getFlagString(cmd, "port")
		id := getFlagString(cmd, "id")
		password := getFlagString(cmd, "password")

		rpcService(
			"RpcService.Join",
			&params.JoinArgs{
				IP:           ip,
				Port:         port,
				RoomId:       id,
				RoomPassword: password,
			})
	},
}

func init() {
	rootCmd.AddCommand(joinCmd)
	joinCmd.Flags().String("ip", "127.0.0.1", "Ip of the websocket")
	joinCmd.Flags().String("port", "8080", "Port")
	joinCmd.Flags().String("id", "defaultRoom", "Id of the room")
	joinCmd.Flags().String("password", "defaultPassword", "Password of the room")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// joinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// joinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
