package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vmmgr/gclient/data"
	"log"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "user",
	Long: `user command. For example:
user add test
user remove test`,
}

var userAddCmd = &cobra.Command{
	Use:   "add",
	Short: "user add",
	Long: `user add tool
for example:
user add [user] [pass]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			log.Fatal("Syntax error!!")
		}

		data.AddUser(cmd, args)
		return nil
	},
}

var userDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "user delete",
	Long: `user add tool
for example:
user delete [UserID]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 0 {
			log.Fatal("Syntax error!!")
		}
		data.DeleteUser(cmd, args)
		return nil
	},
}

var userGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get tool for user",
	Long:  "get tool for user",
}

var userGetAllCmd = &cobra.Command{
	Use:   "all",
	Short: "user get all",
	Long: `get all user
for example:
user get all`,
	RunE: func(cmd *cobra.Command, args []string) error {
		data.GetAllUser(cmd, args)
		return nil
	},
}

var userChangeCmd = &cobra.Command{
	Use:   "change",
	Short: "change tool for user",
	Long:  "change tool for user",
}

var userPassChangeCmd = &cobra.Command{
	Use:   "pass",
	Short: "change pass",
	Long: `change pass tool for user
for example:
user change pass [userID] [newpass]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			log.Fatal("Syntax error!!")
		}
		data.UserNameChange(cmd, args)
		return nil
	},
}

var userNameChangeCmd = &cobra.Command{
	Use:   "name",
	Short: "change name",
	Long: `change name tool for user
for example:
user change name [UserID] [after username]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			log.Fatal("Syntax error!!")
		}

		data.UserPassChange(cmd, args)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
	userCmd.AddCommand(userAddCmd)
	userCmd.AddCommand(userDeleteCmd)
	userCmd.AddCommand(userGetCmd)
	userCmd.AddCommand(userChangeCmd)

	userGetCmd.AddCommand(userGetAllCmd)
	userChangeCmd.AddCommand(userNameChangeCmd)
	userChangeCmd.AddCommand(userPassChangeCmd)
}
