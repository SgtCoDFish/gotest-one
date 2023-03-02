package app

import "github.com/sgtcodfish/gotest-one/caser"

func Transformer(s string) string {
	return caser.Dutch(s)
}
