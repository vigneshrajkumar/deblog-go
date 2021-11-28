package api

type Conf struct {
	Module       string   `yaml:"module"`
	BeforeCreate []string `yaml:"beforeCreate"`
	AfterCreate  []string `yaml:"afterCreate"`
}
