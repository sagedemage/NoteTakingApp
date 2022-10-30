package tests

import (
	"os"
	"path"
	"runtime"

	"github.com/gin-gonic/gin"

	"notebook_app/cmd/app/notebook_db"
	"notebook_app/cmd/app/router_setup"
)

func change_to_repo_root_Dir() {
	/* change to the root of the repo directory */
	_, filename, _, _ := runtime.Caller(0)

  	dir := path.Join(path.Dir(filename), "../../..")
  	err := os.Chdir(dir)
  	if err != nil {
  	  panic(err)
  	}
}

func run_app() *gin.Engine{
	// Open database
	db := notebook_db.UseSQLite("database/notebook.db")

	// setup routerr
	r := router_setup.InitializeRouter(db)

	return r
}

func Setup() *gin.Engine {
	// setup for a test
	change_to_repo_root_Dir() // change directory

	r := run_app() // run app

	return r
}


