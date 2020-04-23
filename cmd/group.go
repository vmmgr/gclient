package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vmmgr/gclient/data"
	"log"
)

// groupCmd represents the group command
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "group command",
	Long: `group command
For example:
group add ..
group delete ..
group get ..
group setting ..
group join ..`,
}

var groupAddCmd = &cobra.Command{
	Use:   "add",
	Short: "group add",
	Long: `group add tool
for example:
Mode: mode0 <- Administrator | mode1 <- User | mode2 <- Team
group add [GroupName] [Mode] -v [MaxVM] -c [MaxCPU] -m [MaxMem] -s [MaxStorage]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			log.Fatal("Syntax error!!")
		}
		data.AddGroup(cmd, args, Spec(cmd))
		return nil
	},
}

var groupDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "group delete",
	Long: `group delete tool
for example:

group delete [GroupID]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			log.Fatal("Syntax error!!")
		}
		data.DeleteGroup(cmd, args)
		return nil
	},
}

var groupGetCmd = &cobra.Command{
	Use:   "get",
	Short: "group remove",
	Long: `group remove tool
for example:`,
}

var groupGetAllCmd = &cobra.Command{
	Use:   "all",
	Short: "group get all",
	Long: `group get tool
for example:

group get all`,
	RunE: func(cmd *cobra.Command, args []string) error {
		data.GetAllGroup(cmd, args)
		return nil
	},
}

var groupGetMyselfCmd = &cobra.Command{
	Use:   "myself",
	Short: "display for group belong to myself",
	Long: `group get tool
for example:
group get myself`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			log.Fatal("Syntax error!!")
		}
		data.GetGroup(cmd, args)
		return nil
	},
}

var groupJoinCmd = &cobra.Command{
	Use:   "join",
	Short: "group join",
	Long: `group join tool
for example:
group join add ..
group join delete ..`,
}

var groupJoinAddCmd = &cobra.Command{
	Use:   "add",
	Short: "group join add",
	Long: `group join add tool
for example:
Mode mode0:Admin | mode1: User
group join add [Mode] [GroupID] [UserID]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			log.Fatal("Syntax error!!")
		}
		data.JoinAddGroup(cmd, args)
		return nil
	},
}

var groupJoinDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "group join delete",
	Long: `group join delete tool
for example:
Mode mode0:Admin | mode1: User
group join delete [Mode] [GroupID] [UserID]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			log.Fatal("Syntax error!!")
		}
		data.JoinDeleteGroup(cmd, args)
		return nil
	},
}

var groupChangeCmd = &cobra.Command{
	Use:   "change",
	Short: "change tool for group",
	Long: `change tool for group
for example:
group change name ..
group change spec ..`,
}

var groupNameChangeCmd = &cobra.Command{
	Use:   "name",
	Short: "change name",
	Long: `change name tool for group
for example:
group change name [GroupID] [new groupname]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			log.Fatal("Syntax error!!")
		}

		data.GroupNameChange(cmd, args)
		return nil
	},
}

var groupSpecChangeCmd = &cobra.Command{
	Use:   "spec",
	Short: "change spec",
	Long: `change spec tool for group
for example:
user change spec [GroupID] -v [MaxVM] -c [MaxCPU] -m [MaxMem] -s [MaxStorage]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 0 {
			log.Fatal("Syntax error!!")
		}
		data.GroupSpecChange(cmd, args, Spec(cmd))
		return nil
	},
}

func Spec(cmd *cobra.Command) data.MaxSpec {
	vm, err := cmd.PersistentFlags().GetInt32("maxvm")
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	cpu, err := cmd.PersistentFlags().GetInt32("maxcpu")
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	mem, err := cmd.PersistentFlags().GetInt32("maxmem")
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	storage, err := cmd.PersistentFlags().GetInt64("maxstorage")
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	return data.MaxSpec{
		MaxVM:      vm,
		MaxCPU:     cpu,
		MaxMem:     mem,
		MaxStorage: storage,
	}
}

func init() {
	rootCmd.AddCommand(groupCmd)
	groupCmd.AddCommand(groupAddCmd)
	groupCmd.AddCommand(groupDeleteCmd)
	groupCmd.AddCommand(groupGetCmd)
	groupGetCmd.AddCommand(groupGetAllCmd)
	groupGetCmd.AddCommand(groupGetMyselfCmd)
	groupCmd.AddCommand(groupJoinCmd)
	groupJoinCmd.AddCommand(groupJoinAddCmd)
	groupJoinCmd.AddCommand(groupJoinDeleteCmd)
	groupCmd.AddCommand(groupChangeCmd)
	groupChangeCmd.AddCommand(groupNameChangeCmd)
	groupChangeCmd.AddCommand(groupSpecChangeCmd)

	groupCmd.PersistentFlags().Int32P("maxvm", "v", 1, "max vm")
	groupCmd.PersistentFlags().Int32P("maxcpu", "c", 1, "max cpu")
	groupCmd.PersistentFlags().Int32P("maxmem", "m", 1024, "max memory")
	groupCmd.PersistentFlags().StringP("maxstorage", "s", "", "max storage capacity")
	groupCmd.PersistentFlags().StringP("net", "n", "", "net")
}
