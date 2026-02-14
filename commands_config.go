package main

type commandsConfig struct {
	locationArea paginationConfig
}

type paginationConfig struct {
	next     string
	previous string
}

func newCommandsConfig() *commandsConfig {
	return &commandsConfig{
		locationArea: paginationConfig{
			next:     "",
			previous: "",
		},
	}
}
