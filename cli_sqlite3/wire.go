//go:build wireinject

package cli_sqlite3

import "github.com/google/wire"

func InitApp() (*App, error) {
	panic(wire.Build(wires))
}
