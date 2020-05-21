package cmd

import (
	"github.com/gdejong/adb-golang/pkg/adb"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(playStickHeroCommand)
}

var playStickHeroCommand = &cobra.Command{
	Use:   "adb:play-stick-hero",
	Short: "Play the Stick Hero game",
	Run: func(cmd *cobra.Command, args []string) {
		adb.MakeScreenshot("./screenie.png")
		// TODO finish me, for now it just makes a screenshot
	},
}
