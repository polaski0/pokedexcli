package main

import "os"

func commandExit(c *Config, args ...string) error {
	os.Exit(0)
	return nil
}
