package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vmmgr/gclient/data"
	"log"
)

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "create: token create ,delete: token delete",
	Long: `For example:
`,
}
var tokenGenerateCmd = &cobra.Command{
	Use: "generate",
	Short: `token generate 
0:Permanent 1:24H 2:1H`,
	Long: `token generate tool
for example:
0:Permanent 1:24H 2:1H
token generate　[0/1/2]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		user, err := cmd.Flags().GetString("user")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		pass, err := cmd.Flags().GetString("pass")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		data.GenerateToken(cmd, args, user, pass)
		return nil
	},
}

var tokenDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "token delete",
	Long: `token delete tool
for example:

token delete [token]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 || 2 < len(args) {
			log.Fatal("Syntax error!!")
		}
		data.DeleteToken(cmd, args)
		return nil
	},
}

var tokenGetAllCmd = &cobra.Command{
	Use:   "get",
	Short: "get tokens for all users",
	Long: `token get tool
for example:

token get`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			log.Fatal("Syntax error!!")
		}
		data.GetAllToken(cmd, args)
		return nil
	},
}

func init() {

	rootCmd.AddCommand(tokenCmd)
	tokenCmd.AddCommand(tokenGenerateCmd)
	tokenCmd.AddCommand(tokenDeleteCmd)
	tokenCmd.AddCommand(tokenGetAllCmd)

	tokenCmd.PersistentFlags().StringP("user", "u", "", "username")
	tokenCmd.PersistentFlags().StringP("pass", "p", "", "password")
}
