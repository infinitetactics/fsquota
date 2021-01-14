package main

import (
	"os/user"

	"github.com/spf13/cobra"
)

var cmdGroup = &cobra.Command{
	Use:   "group",
	Short: "Group quota management",
}

func init() {
	cmdRoot.AddCommand(cmdGroup)
}

func lookupGroup(groupIDOrGroupName string) (grp *user.Group, err error) {
	if isNumeric(groupIDOrGroupName) {
		grp = &user.Group{
			Gid: groupIDOrGroupName,
		}
		return
	}
	return user.LookupGroup(groupIDOrGroupName)
}
