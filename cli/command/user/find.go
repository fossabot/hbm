package user

import (
	"fmt"

	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/cli/command"
	"github.com/kassisol/hbm/storage"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func newFindCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "find [name]",
		Short: "Verify if user exists in the whitelist",
		Long:  findDescription,
		Args:  cobra.ExactArgs(1),
		Run:   runFind,
	}

	return cmd
}

func runFind(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", command.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer s.End()

	result := s.FindUser(args[0])

	fmt.Println(result)
}

var findDescription = `
Verify if user exists in the whitelist

`
