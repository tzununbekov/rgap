package main

import (
	"fmt"
	"net"
	"time"

	"github.com/spf13/cobra"

	"github.com/Snawoot/rgap"
)

var (
	group    uint64
	address  net.IP
	key      pskOption
	interval time.Duration
)

type pskOption struct {
	psk *rgap.PSK
}

func (psk pskOption) String() string {
	if psk.psk == nil {
		return "<nil>"
	}
	return psk.psk.String()
}

func (psk pskOption) Set(s string) error {
	return psk.psk.FromHexString(s)
}

func (psk pskOption) Type() string {
	return "hexstring"
}

// agentCmd represents the agent command
var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Run agent to send announcements",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(group, address, key.psk, interval)
	},
}

func init() {
	rootCmd.AddCommand(agentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// agentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// agentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	agentCmd.Flags().Uint64VarP(&group, "group", "g", 0, "redundancy group")
	agentCmd.Flags().IPVarP(&address, "address", "a", nil, "IP address to announce")
	agentCmd.Flags().VarP(&key, "psk", "k", "pre-shared key for announcement signature")
}
