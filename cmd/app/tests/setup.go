package tests

import (
	"os"
	"path"
	"runtime"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"notebook_app/cmd/app/notebook_db"
	"notebook_app/cmd/app/router_setup"
)

func ChangetoRepoRootDirectory() {
	_, filename, _, _ := runtime.Caller(0)

  	dir := path.Join(path.Dir(filename), "../../..")
  	err := os.Chdir(dir)
  	if err != nil {
  	  panic(err)
  	}
}

func RunApp() (*gin.Engine, *gorm.DB){
	// Open database
	db := notebook_db.InitDB("database/notebook.db")

	// setup routerr
	r := router_setup.InitializeRouter(db)

	return r, db
}


