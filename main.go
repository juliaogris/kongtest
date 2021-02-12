package main

import (
	"fmt"

	"github.com/alecthomas/kong"
)

type config struct {
	Arg string `help: "required argument"`
}

func (cfg *config) Validate() error {
	if cfg.Arg == "" {
		return fmt.Errorf("missing Arg")
	}
	return nil
}

func main() {
	cfg := &config{}
	_ = kong.Parse(cfg)
	fmt.Println("hello")
}
