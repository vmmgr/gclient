package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vmmgr/gclient/data"
	"log"
)

var nicCmd = &cobra.Command{
	Use:   "nic",
	Short: "nic",
	Long: `nic command. For example:
nic add test
nic delete test`,
}

var nicAddCmd = &cobra.Command{
	Use:   "add",
	Short: "nic add -n [Name] -g [GroupID] -N [NetID] -d [Driver]",
	Long: `nic add tool
for example:
Driver | 1:virtio
nic add -n [Name] -g [GroupID] -N [NetID] -d [Driver]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			log.Fatal("Syntax error!!")
		}

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		gid, err := cmd.Flags().GetInt64("gid")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		nid, err := cmd.Flags().GetInt64("nid")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		driver, err := cmd.Flags().GetInt32("driver")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}

		data.AddNIC(cmd, data.NICData{Name: name, Gid: gid, Nid: nid, Driver: driver})
		return nil
	},
}

var nicDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "nic delete",
	Long: `nic add tool
for example:
nic delete [NIC ID]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 0 {
			log.Fatal("Syntax error!!")
		}
		data.DeleteNIC(cmd, args)
		return nil
	},
}

var nicGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get tool for nic",
	Long:  "get tool for nic",
}

var nicGetAllCmd = &cobra.Command{
	Use:   "all",
	Short: "nic get all",
	Long: `get all nic
for example:
nic get all`,
	RunE: func(cmd *cobra.Command, args []string) error {
		data.GetAllNIC(cmd, args)
		return nil
	},
}

var nicChangeCmd = &cobra.Command{
	Use:   "change",
	Short: "change tool for nic",
	Long:  "change tool for nic",
}

var nicNameChangeCmd = &cobra.Command{
	Use:   "name",
	Short: "change name",
	Long: `change name tool for nic
for example:
nic change name [NICID] [after nicname]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			log.Fatal("Syntax error!!")
		}

		//data.NICNameChange(cmd, args)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(nicCmd)
	nicCmd.AddCommand(nicAddCmd)
	nicCmd.AddCommand(nicDeleteCmd)
	nicCmd.AddCommand(nicGetCmd)
	nicCmd.AddCommand(nicChangeCmd)

	nicGetCmd.AddCommand(nicGetAllCmd)
	nicChangeCmd.AddCommand(nicNameChangeCmd)

	nicAddCmd.PersistentFlags().StringP("name", "n", "", "name")
	nicAddCmd.PersistentFlags().Int64P("gid", "g", 0, "gID")
	nicAddCmd.PersistentFlags().Int64P("nid", "N", 0, "nID")
	nicAddCmd.PersistentFlags().Int32P("driver", "d", 0, "driver")
}
