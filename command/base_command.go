package command

type BaseCommand interface {
	CommandName() string
	Exec(ctx map[string]interface{}) error
}
