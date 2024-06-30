package seeder

import (
	"encoding/json"
	"os"

	"github.com/BeepLoop/nearbyassist_seeder/config"
	"github.com/BeepLoop/nearbyassist_seeder/database"
	"github.com/BeepLoop/nearbyassist_seeder/request"
	"golang.org/x/crypto/bcrypt"
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

				if _, err := s.Db.InsertTag(req); err != nil {
					someErr = err
					break
				}
			}
		case "admin":
			for _, admin := range entry.TableData {
				req := &request.AdminModel{}
				if err := s.JsonRawToStruct(admin, req); err != nil {
					someErr = err
					break
				}

				hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
				if err != nil {
					someErr = err
					break
				} else {
					req.Password = string(hash)
				}

				if _, err := s.Db.InsertAdmin(req); err != nil {
					someErr = err
					break
				}
			}
		case "user":
			for _, user := range entry.TableData {
				req := &request.UserModel{}
				if err := s.JsonRawToStruct(user, req); err != nil {
					someErr = err
					break
				}

				if _, err := s.Db.InsertUser(req); err != nil {
					someErr = err
					break
				}
			}
		case "vendor":
			for _, vendor := range entry.TableData {
				req := &request.VendorModel{}
				if err := s.JsonRawToStruct(vendor, req); err != nil {
					someErr = err
					break
				}

				if _, err := s.Db.InsertVendor(req); err != nil {
					someErr = err
					break
				}
			}
		case "service":
			services := make([]*request.ServiceModel, 0)
			for _, service := range entry.TableData {
				req := &request.ServiceModel{}
				if err := s.JsonRawToStruct(service, req); err != nil {
					someErr = err
					break
				}

				services = append(services, req)
			}

			for _, service := range services {
				id, err := s.Db.InsertService(service)
				if err != nil {
					someErr = err
					break
				}

				for _, tag := range service.Tags {
					svcTag := &request.ServiceTagModel{
						ServiceId: id,
						TagTitle:  tag.Title,
					}
					if _, err := s.Db.InsertServiceTag(svcTag); err != nil {
						someErr = err
						break
					}
				}
			}
		case "review":
			for _, review := range entry.TableData {
				req := &request.ReviewModel{}
				if err := s.JsonRawToStruct(review, req); err != nil {
					someErr = err
					break
				}

				if _, err := s.Db.InsertReview(req); err != nil {
					someErr = err
					break
				}
			}
		case "servicePhoto":
			for _, servicePhoto := range entry.TableData {
				req := &request.ServicePhotoModel{}
				if err := s.JsonRawToStruct(servicePhoto, req); err != nil {
					someErr = err
					break
				}

				if _, err := s.Db.InsertServicePhoto(req); err != nil {
					someErr = err
					break
				}
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
