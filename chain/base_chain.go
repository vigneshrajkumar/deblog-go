package chain

import (
	"deblog-go/command"
	"log"
)

type BaseChain struct {
	commands []command.BaseCommand
}

func (bc *BaseChain) AddCommand(cmd command.BaseCommand) {
	bc.commands = append(bc.commands, cmd)
}

func (bc *BaseChain) Exec(ctx map[string]interface{}) error {
	for ix, cmd := range bc.commands {
		log.Println("executing cmd ", ix, " ", cmd.CommandName())
		if err := cmd.Exec(ctx); err != nil {
			return err
		}
	}
	return nil
}
