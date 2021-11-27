package chain

import "deblog-go/command"

type Chain interface {
	AddCommand(cmd command.BaseCommand)
	Exec(ctx map[string]interface{}) error
}
