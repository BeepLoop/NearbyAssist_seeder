package main

import (
	"bufio"
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
	key := flag.String("key", "N1PCdw3M2B1TfJhoaY2mL736p2vCUc47", "AES encryption key")
	verbose := flag.Bool("v", false, "Verbose mode")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)

	if *headless {
		fmt.Println("Running in headless mode...")

		// Load Config
		conf := config.NewEmptyConfig()
		conf.LoadConfigFromEnv()

		// Load Database
		db := database.NewMysqlDatabase(conf)
		if err := db.InitConnection(); err != nil {
			fmt.Println("Failed to connect to database, error: " + err.Error())
			scanner.Scan()
			os.Exit(1)
		}

		// Load Seeder backend
		seeder := seeder.NewSeeder(conf, db, *key, *verbose)
		if err := seeder.Seed(); err != nil {
			fmt.Println("Failed to seed data, error: " + err.Error())
			scanner.Scan()
			os.Exit(1)
		}

		fmt.Println("Data seeded successfully! Click Enter to exit...")
		scanner.Scan()
		os.Exit(0)
	}
}
