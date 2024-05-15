package main

import (
	cli "github.com/hayeah/cli_sqlite3"
)

func main() {
	app, err := cli.InitApp()
	if err != nil {
		panic(err)
	}

	err = app.Run()
	if err != nil {
		panic(err)
	}
}
