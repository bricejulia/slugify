package command

import (
	"fmt"
	"github.com/bricejulia/slugify/repl"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/user"
)

var rootCmd = &cobra.Command{
	Use: "slugify",
	Run: func(cmd *cobra.Command, args []string) {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Hello %s!\n", user.Username)
		fmt.Printf("Feel free to type text to slugify\n")
		repl.Start(os.Stdin, os.Stdout)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
