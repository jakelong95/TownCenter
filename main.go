package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/jakelong95/TownCenter/router"

	"github.com/ghmeier/bloodlines/config"
)

func main() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	/* Load configuration */
	config, err := config.Init(path.Join(dir, "config.json"))
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	tc, err := router.New(config)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	fmt.Printf("TownCenter is now running on %s\n", config.Port)
	tc.Start(":" + config.Port)
}
