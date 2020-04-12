package integrationtests

import (
	"fmt"
	"os"
	"testing"

	"github.com/akantsevoi/integrationtests/cmd/database"
	"github.com/akantsevoi/integrationtests/cmd/server"
)

func TestMain(m *testing.M) {

	options := database.Options{
		Name:     "root",
		Password: "test",
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBname:   "Integration",
	}
	s, err := server.NewServer(options)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	router = s.Router

	code := m.Run()
	os.Exit(code)
}
