package app

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

type App struct {
	db *sql.DB
}

func NewApp() *App {
	return &App {

	}
}

func (a *App) Start() {
	server := gin.Default()

}