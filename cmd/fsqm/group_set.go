package main

import (
	"errors"
	"os/user"

	"github.com/infinitetactics/fsquota"
	"github.com/speijnik/go-errortree"
	"github.com/spf13/cobra"
)

var cmdGroupSet = &cobra.Command{
	Use:   "set path group",
	Short: "Sets quota configuration for a given user",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if len(args) != 2 {
			err = errors.New("exactly two arguments required")
			return
		}

		var bytesSoft, bytesHard, filesSoft, filesHard uint64
		var bytesPresent, filesPresent bool
		var parseErr error

		if bytesSoft, bytesHard, bytesPresent, parseErr = parseLimitsFlag(cmd, "bytes"); parseErr != nil {
			err = errortree.Add(err, "bytes", parseErr)
		}

		if filesSoft, filesHard, filesPresent, parseErr = parseLimitsFlag(cmd, "files"); parseErr != nil {
			err = errortree.Add(err, "files", parseErr)
		}

		if err != nil {
			return
		}

		var g *user.Group
		if g, err = lookupGroup(args[1]); err != nil {
			return
		}

		var info *fsquota.Info
		limits := fsquota.Limits{}

		if bytesPresent {
			limits.Bytes.SetSoft(bytesSoft)
			limits.Bytes.SetHard(bytesHard)
		}

		if filesPresent {
			limits.Files.SetSoft(filesSoft)
			limits.Files.SetHard(filesHard)
		}

		if !filesPresent && !bytesPresent {
			err = errors.New("nothing to set")
			return
		}

		if info, err = fsquota.SetGroupQuota(args[0], g, &limits); err != nil {
			return
		}

		printQuotaInfo(cmd, info)
		return
	},
}

func init() {
	cmdGroupSet.Flags().StringP("bytes", "b", "", "Byte limit in soft,hard format. ie. 1MiB,2GiB")
	cmdGroupSet.Flags().StringP("files", "f", "", "File limit in soft,hard format, ie. 1M,2G")
	cmdGroup.AddCommand(cmdGroupSet)
}
