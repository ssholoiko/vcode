package cmd

import (
	"fmt"
	"net"
	"time"

	"github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
	Use:   "ping [host]",
	Short: "Пінгує вказаний хост через TCP:80",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		host := args[0]
		conn, err := net.DialTimeout("tcp", host+":80", 2*time.Second)
		if err != nil {
			fmt.Printf("❌ %s не доступний: %s\n", host, err)
			return
		}
		conn.Close()
		fmt.Printf("✅ %s доступний (порт 80)\n", host)
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)
}
