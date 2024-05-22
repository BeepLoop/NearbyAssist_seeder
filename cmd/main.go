package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/BeepLoop/nearbyassist_seeder/config"
	"github.com/BeepLoop/nearbyassist_seeder/database"
	"github.com/BeepLoop/nearbyassist_seeder/seeder"
)

const (
	IS_HEADLESS = true
)

func main() {
	headless := flag.Bool("headless", IS_HEADLESS, "Run without GUI. Defaults to true")
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
			fmt.Println("Failed to seed data, error: " + err.Error())
			os.Exit(1)
		}

		fmt.Println("Data seeded successfully!")
		os.Exit(0)
	}
}
