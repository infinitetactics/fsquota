package main

import (
	"os/user"

	"github.com/spf13/cobra"
)

var cmdUser = &cobra.Command{
	Use:   "user",
	Short: "User quota management",
}

func init() {
	cmdRoot.AddCommand(cmdUser)
}

func lookupUser(userIDOrUsername string) (usr *user.User, err error) {
	if isNumeric(userIDOrUsername) {
		usr = &user.User{
			Uid: userIDOrUsername,
		}
		return
	}
	return user.Lookup(userIDOrUsername)
}
