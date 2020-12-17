package main

import "github.com/HunterGooD/testsGolang/Authorization/app"

func main() {
	app := app.NewApp()
	app.Init("db.sqlite3")
	app.Start(":8080")
}
