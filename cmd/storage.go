package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vmmgr/gclient/data"
	"log"
)

var storageCmd = &cobra.Command{
	Use:   "storage",
	Short: "storage",
	Long: `storage command. For example:
storage add test
storage delete test`,
}

var storageAddCmd = &cobra.Command{
	Use:   "add",
	Short: "storage add [Name] [GroupID] [Driver] [Size(MB)] [Mode] [Path] [Image]",
	Long: `storage add tool
for example:
Driver| 1:virtio
Mode  | 1~9:AutoMode 10:ManualMode
storage add [Name] [GroupID] [Driver] [Size(MB)] [Mode] [Path] [Image]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			log.Fatal("Syntax error!!")
		}

		data.AddStorage(cmd, args)
		return nil
	},
}

var storageDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "storage delete",
	Long: `storage add tool
for example:
storage delete [StorageID]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 0 {
			log.Fatal("Syntax error!!")
		}
		//data.DeleteStorage(cmd, args)
		return nil
	},
}

var storageGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get tool for storage",
	Long:  "get tool for storage",
}

var storageGetAllCmd = &cobra.Command{
	Use:   "all",
	Short: "storage get all",
	Long: `get all storage
for example:
storage get all`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//data.GetAllStorage(cmd, args)
		return nil
	},
}

var storageChangeCmd = &cobra.Command{
	Use:   "change",
	Short: "change tool for storage",
	Long:  "change tool for storage",
}

var storageNameChangeCmd = &cobra.Command{
	Use:   "name",
	Short: "change name",
	Long: `change name tool for storage
for example:
storage change name [StorageID] [after storagename]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			log.Fatal("Syntax error!!")
		}

		//data.StorageNameChange(cmd, args)
		return nil
	},
}

var storagePassChangeCmd = &cobra.Command{
	Use:   "pass",
	Short: "change pass",
	Long: `change pass tool for storage
for example:
storage change pass [storageID] [newpass]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			log.Fatal("Syntax error!!")
		}
		//data.StoragePassChange(cmd, args)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(storageCmd)
	storageCmd.AddCommand(storageAddCmd)
	storageCmd.AddCommand(storageDeleteCmd)
	storageCmd.AddCommand(storageGetCmd)
	storageCmd.AddCommand(storageChangeCmd)

	storageGetCmd.AddCommand(storageGetAllCmd)
	storageChangeCmd.AddCommand(storageNameChangeCmd)
	storageChangeCmd.AddCommand(storagePassChangeCmd)
}
