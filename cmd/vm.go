package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vmmgr/gclient/data"
	"log"
)

var vmCmd = &cobra.Command{
	Use:   "vm",
	Short: "vm",
	Long: `vm command. For example:
vm add test
vm delete test`,
}

var vmAddCmd = &cobra.Command{
	Use:   "add",
	Short: "vm add [Name] [GroupID] [Driver] [Size(MB)] [Mode] [Path] [Image]",
	Long: `vm add tool
for example:
Driver| 1:virtio
Mode  | 1~9:AutoMode 10:ManualMode
vm add [Name] [GroupID] [Driver] [Size(MB)] [Mode] [Path] [Image]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			log.Fatal("Syntax error!!")
		}

		data.AddStorage(cmd, args)
		return nil
	},
}

var vmDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "vm delete",
	Long: `vm add tool
for example:
vm delete [vmID]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 0 {
			log.Fatal("Syntax error!!")
		}
		data.DeleteVM(cmd, args)
		return nil
	},
}

//0:Start 1:Stop 2:Shutdown 3:Restart 4:HardReset 5:Pause 6:Resume 10:SnapShot

var vmStartCmd = &cobra.Command{
	Use:   "start",
	Short: "vm start",
	Long: `vm start tool
for example:
vm start [vmID]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 0 {
			log.Fatal("Syntax error!!")
		}
		data.StatusVM(cmd, args, 0)
		return nil
	},
}

var vmStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "vm stop",
	Long: `vm stop tool
for example:
vm stop [vmID]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 0 {
			log.Fatal("Syntax error!!")
		}
		data.StatusVM(cmd, args, 1)
		return nil
	},
}

var vmShutdownCmd = &cobra.Command{
	Use:   "shutdown",
	Short: "vm shutdown",
	Long: `vm shutdown tool
for example:
vm shutdown [vmID]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 0 {
			log.Fatal("Syntax error!!")
		}
		data.StatusVM(cmd, args, 2)
		return nil
	},
}

var vmResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "vm reset",
	Long: `vm reset tool
for example:
vm reset [vmID]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 0 {
			log.Fatal("Syntax error!!")
		}
		data.StatusVM(cmd, args, 4)
		return nil
	},
}

var vmPauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "vm pause",
	Long: `vm pause tool
for example:
vm pause [vmID]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 0 {
			log.Fatal("Syntax error!!")
		}
		data.StatusVM(cmd, args, 5)
		return nil
	},
}

var vmResumeCmd = &cobra.Command{
	Use:   "resume",
	Short: "vm resume",
	Long: `vm resume tool
for example:
vm resume [vmID]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 0 {
			log.Fatal("Syntax error!!")
		}
		data.StatusVM(cmd, args, 6)
		return nil
	},
}

var vmGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get tool for vm",
	Long:  "get tool for vm",
}

var vmGetAllCmd = &cobra.Command{
	Use:   "all",
	Short: "vm get all",
	Long: `get all vm
for example:
vm get all`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//data.GetAllStorage(cmd, args)
		return nil
	},
}

var vmChangeCmd = &cobra.Command{
	Use:   "change",
	Short: "change tool for vm",
	Long:  "change tool for vm",
}

var vmNameChangeCmd = &cobra.Command{
	Use:   "name",
	Short: "change name",
	Long: `change name tool for vm
for example:
vm change name [StorageID] [after vmname]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			log.Fatal("Syntax error!!")
		}

		//data.StorageNameChange(cmd, args)
		return nil
	},
}

var vmPassChangeCmd = &cobra.Command{
	Use:   "pass",
	Short: "change pass",
	Long: `change pass tool for vm
for example:
vm change pass [vmID] [newpass]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			log.Fatal("Syntax error!!")
		}
		//data.StoragePassChange(cmd, args)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(vmCmd)
	vmCmd.AddCommand(vmAddCmd)
	vmCmd.AddCommand(vmDeleteCmd)
	vmCmd.AddCommand(vmStartCmd)
	vmCmd.AddCommand(vmStopCmd)
	vmCmd.AddCommand(vmShutdownCmd)
	vmCmd.AddCommand(vmResetCmd)
	vmCmd.AddCommand(vmPauseCmd)
	vmCmd.AddCommand(vmResumeCmd)
	vmCmd.AddCommand(vmGetCmd)
	vmCmd.AddCommand(vmChangeCmd)

	vmGetCmd.AddCommand(vmGetAllCmd)
	vmChangeCmd.AddCommand(vmNameChangeCmd)
	vmChangeCmd.AddCommand(vmPassChangeCmd)
}
