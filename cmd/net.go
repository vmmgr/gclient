package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vmmgr/gclient/data"
	"log"
)

var netCmd = &cobra.Command{
	Use:   "net",
	Short: "net",
	Long: `net command. For example:
net add test
net delete test`,
}

var netAddCmd = &cobra.Command{
	Use:   "add",
	Short: "net add -n [Name] -g [GroupID] -V [VLAN]",
	Long: `net add tool
for example:
Driver | 1:virtio
net add -n [Name] -g [GroupID] -V [VLAN]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			log.Fatal("Syntax error!!")
		}

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		gid, err := cmd.Flags().GetString("gid")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		vlan, err := cmd.Flags().GetInt32("vlan")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}

		data.AddNet(cmd, data.NetData{Name: name, Gid: gid, Vlan: vlan})
		return nil
	},
}

var netDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "net delete",
	Long: `net add tool
for example:
net delete [Net ID]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 0 {
			log.Fatal("Syntax error!!")
		}
		data.DeleteNet(cmd, args)
		return nil
	},
}

var netGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get tool for net",
	Long:  "get tool for net",
}

var netGetAllCmd = &cobra.Command{
	Use:   "all",
	Short: "net get all",
	Long: `get all net
for example:
net get all`,
	RunE: func(cmd *cobra.Command, args []string) error {
		data.GetAllNet(cmd, args)
		return nil
	},
}

var netChangeCmd = &cobra.Command{
	Use:   "change",
	Short: "change tool for net",
	Long:  "change tool for net",
}

var netNameChangeCmd = &cobra.Command{
	Use:   "name",
	Short: "change name",
	Long: `change name tool for net
for example:
net change name [NetID] [after netname]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			log.Fatal("Syntax error!!")
		}

		//data.NetNameChange(cmd, args)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(netCmd)
	netCmd.AddCommand(netAddCmd)
	netCmd.AddCommand(netDeleteCmd)
	netCmd.AddCommand(netGetCmd)
	netCmd.AddCommand(netChangeCmd)

	netGetCmd.AddCommand(netGetAllCmd)
	netChangeCmd.AddCommand(netNameChangeCmd)

	netAddCmd.PersistentFlags().StringP("name", "n", "", "name")
	netAddCmd.PersistentFlags().StringP("gid", "g", 0, "gID")
	netAddCmd.PersistentFlags().Int32P("vlan", "v", 0, "vlan")
}
