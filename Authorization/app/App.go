package app

import (
	"database/sql"

	"github.com/HunterGooD/testsGolang/Authorization/db"
	"github.com/HunterGooD/testsGolang/Authorization/utils"

	"github.com/gin-gonic/gin"
)

// App главная структура прложения
type App struct {
	db *sql.DB
}

// NewApp конструктор структуры App
func NewApp() *App {
	return &App{}
}

// Init инициализация прилложения и БД
func (a *App) Init(dbname string) {
	var (
		Db  *sql.DB
		err error
	)

	if !utils.FileIsExist(dbname) {
		Db, err = db.CreateDB(dbname)
		if err != nil {
			panic(err)
		}
	} else {
		Db, err = db.ConnectDB(dbname)
	}

	a.db = Db
}

// Start стартует приложение
func (a *App) Start(addr string) {
	server := gin.Default()
	server.POST("/signup", a.signUp)
	server.POST("/signin", a.signIn)
	server.Run(addr)
}
