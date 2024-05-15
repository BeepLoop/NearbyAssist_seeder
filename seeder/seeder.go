package seeder

import (
	"encoding/json"
	"fmt"
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

	var someErr error
	for _, entry := range data.Data {
		switch entry.Table {
		case "tag":
			for _, tag := range entry.TableData {
				req := &request.TagModel{}
				if err := s.JsonRawToStruct(tag, req); err != nil {
					someErr = err
					break
				}

				// Do something about req
				fmt.Println(req)
			}
		case "admin":
			for _, admin := range entry.TableData {
				req := &request.AdminModel{}
				if err := s.JsonRawToStruct(admin, req); err != nil {
					someErr = err
					break
				}

				// Do something about req
				fmt.Println(req)
			}
		case "user":
			for _, user := range entry.TableData {
				req := &request.UserModel{}
				if err := s.JsonRawToStruct(user, req); err != nil {
					someErr = err
					break
				}

				// Do something about req
				fmt.Println(req)
			}
		case "vendor":
			for _, vendor := range entry.TableData {
				req := &request.VendorModel{}
				if err := s.JsonRawToStruct(vendor, req); err != nil {
					someErr = err
					break
				}

				// Do something about req
				fmt.Println(req)
			}
		case "service":
			for _, service := range entry.TableData {
				req := &request.ServiceModel{}
				if err := s.JsonRawToStruct(service, req); err != nil {
					someErr = err
					break
				}

				// Do something about req
				fmt.Println(req)
			}
		case "review":
			for _, review := range entry.TableData {
				req := &request.ReviewModel{}
				if err := s.JsonRawToStruct(review, req); err != nil {
					someErr = err
					break
				}

				// Do something about req
				fmt.Println(req)
			}
		case "servicePhoto":
			for _, servicePhoto := range entry.TableData {
				req := &request.ServicePhotoModel{}
				if err := s.JsonRawToStruct(servicePhoto, req); err != nil {
					someErr = err
					break
				}

				// Do something about req
				fmt.Println(req)
			}
		default:
			println("unsupported table")
		}
	}

	if someErr != nil {
		return someErr
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

// dist must be a pointer to a struct
func (s *Seeder) JsonRawToStruct(source json.RawMessage, dist any) error {
	bytes, err := source.MarshalJSON()
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, dist); err != nil {
		return err
	}

	return nil
}
