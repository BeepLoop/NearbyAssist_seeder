package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/BeepLoop/nearbyassist_seeder/config"
	"github.com/BeepLoop/nearbyassist_seeder/database"
	"github.com/BeepLoop/nearbyassist_seeder/seeder"
)

func main() {
	headless := flag.Bool("headless", false, "Run without GUI")
	flag.Parse()

	if *headless {
		fmt.Println("Running in headless mode...")

		// Load Config
		conf := config.NewEmptyConfig()
		conf.LoadConfigFromEnv()

		// Load Database
		db := database.NewMysqlDatabase(conf)
		if err := db.InitConnection(); err != nil {
			panic("Failed to connect to database, error: " + err.Error())
		}

		// Load Seeder backend
		seeder := seeder.NewSeeder(conf, db)
		if err := seeder.Seed(); err != nil {
			panic("Failed to seed, error: " + err.Error())
		}

		os.Exit(0)
	}
}
