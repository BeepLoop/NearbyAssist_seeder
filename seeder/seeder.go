package seeder

import (
	"encoding/json"
	"os"

	"github.com/BeepLoop/nearbyassist_seeder/config"
	"github.com/BeepLoop/nearbyassist_seeder/database"
	"github.com/BeepLoop/nearbyassist_seeder/request"
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
	data, err := s.ReadJsonData()
	if err != nil {
		return err
	}

	for _, entry := range data.Data {
		println(entry.Table)
	}

	return nil
}

func (s *Seeder) ReadJsonData() (*request.JsonData, error) {
	bytes, err := os.ReadFile(s.Config.SourceData)
	if err != nil {
		return nil, err
	}

	data := &request.JsonData{}
	if err := json.Unmarshal(bytes, data); err != nil {
		return nil, err
	}

	return data, nil
}
