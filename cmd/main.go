package main

import (
	"github.com/ShekleinAleksey/crm"
)

func main() {
	srv := new(crm.Server)
	if err := srv.Run("8000"); err != nil {

	}
}
