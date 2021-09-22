package config

import (
	"fmt"
	"testing"
)

func TestInitConfig(t *testing.T) {
	initConfig()
	fmt.Printf("%#v\n", Get())
}

func TestValid(t *testing.T) {
	Get().validConfig(validRule)
}
