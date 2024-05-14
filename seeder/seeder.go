package seeder

import (
	"github.com/BeepLoop/nearbyassist_seeder/config"
	"github.com/BeepLoop/nearbyassist_seeder/database"
)

type Seeder struct {
	Config *config.Config
	Db     database.Database
}

func NewSeeder(conf *config.Config, db database.Database) *Seeder {
	return &Seeder{
		Config: conf,
		Db:     db,
	}
}

func (s *Seeder) Seed() error {
	println("Seeding...")
	return nil
}
