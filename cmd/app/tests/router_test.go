package tests

import (
	"os"
  	"path"
  	"runtime"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go-web-app-experiment/cmd/app/router_setup"
)

func TestRoutes(t *testing.T) {
  	_, filename, _, _ := runtime.Caller(0)

  	dir := path.Join(path.Dir(filename), "../../..")
  	err := os.Chdir(dir)
  	if err != nil {
  	  panic(err)
  	}

	r := router_setup.InitializeRouter()

	write := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(write, request)

	assert.Equal(t, 200, write.Code)
}
