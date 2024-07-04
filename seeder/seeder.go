package seeder

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	"github.com/BeepLoop/nearbyassist_seeder/config"
	"github.com/BeepLoop/nearbyassist_seeder/database"
	"github.com/BeepLoop/nearbyassist_seeder/request"
	"github.com/BeepLoop/nearbyassist_seeder/utils"
	"github.com/beeploop/aes-encrypt/encrypt"
	"golang.org/x/crypto/bcrypt"
)

type Seeder struct {
	Config  *config.Config
	Db      database.Database
	Key     string
	verbose bool
}

func NewSeeder(conf *config.Config, db database.Database, key string, verbose bool) *Seeder {
	return &Seeder{
		Config:  conf,
		Db:      db,
		Key:     key,
		verbose: verbose,
	}
}

func (s *Seeder) Seed() error {
	data, err := s.ReadJsonData()
	if err != nil {
		return err
	}

	crypto, err := encrypt.New(s.Key)
	if err != nil {
		return err
	}

	for _, entry := range data.Data {
		switch entry.Table {
		case "tag":
			for _, tag := range entry.TableData {
				req := &request.TagModel{}
				if err := s.JsonRawToStruct(tag, req); err != nil {
					return err
				}

				if id, err := s.Db.InsertTag(req); err != nil {
					return err
				} else {
					if s.verbose {
						fmt.Println("tag: ", req.Title, "; id: ", id)
					}
				}
			}
		case "admin":
			for _, admin := range entry.TableData {
				req := &request.AdminModel{}
				if err := s.JsonRawToStruct(admin, req); err != nil {
					return err
				}

				if hash, err := utils.Hash(req.Username); err != nil {
					return err
				} else {
					req.UsernameHash = hash
				}

				if cipher, err := crypto.Encrypt([]byte(req.Username)); err != nil {
					return err
				} else {
					req.Username = hex.EncodeToString(cipher)
				}

				hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
				if err != nil {
					return err
				} else {
					req.Password = string(hash)
				}

				if id, err := s.Db.InsertAdmin(req); err != nil {
					return err
				} else {
					if s.verbose {
						fmt.Println("admin: ", req.Username, "; id: ", id)
					}
				}
			}
		case "user":
			for _, user := range entry.TableData {
				req := &request.UserModel{}
				if err := s.JsonRawToStruct(user, req); err != nil {
					return err
				}

				if hash, err := utils.Hash(req.Email); err != nil {
					return err
				} else {
					req.EmailHash = hash
				}

				if cipher, err := crypto.Encrypt([]byte(req.Name)); err != nil {
					return err
				} else {
					req.Name = hex.EncodeToString(cipher)
				}

				if cipher, err := crypto.Encrypt([]byte(req.Email)); err != nil {
					return err
				} else {
					req.Email = hex.EncodeToString(cipher)
				}

				if id, err := s.Db.InsertUser(req); err != nil {
					return err
				} else {
					if s.verbose {
						fmt.Println("user: ", req.Email, "; id: ", id)
					}
				}
			}
		case "vendor":
			for _, vendor := range entry.TableData {
				req := &request.VendorModel{}
				if err := s.JsonRawToStruct(vendor, req); err != nil {
					return err
				}

				if hashed, err := utils.Hash(req.Email); err != nil {
					return err
				} else {
					req.Email = hashed
				}

				if id, err := s.Db.InsertVendor(req); err != nil {
					return err
				} else {
					if s.verbose {
						fmt.Println("vendor: ", req.Email, "; id: ", id)
					}
				}
			}
		case "service":
			services := make([]*request.ServiceModel, 0)
			for _, service := range entry.TableData {
				req := &request.ServiceModel{}
				if err := s.JsonRawToStruct(service, req); err != nil {
					return err
				}

				services = append(services, req)
			}

			for _, service := range services {
				if hashed, err := utils.Hash(service.Vendor); err != nil {
					return err
				} else {
					service.Vendor = hashed
				}

				id, err := s.Db.InsertService(service)
				if err != nil {
					return err
				}

				for _, tag := range service.Tags {
					svcTag := &request.ServiceTagModel{
						ServiceId: id,
						TagTitle:  tag.Title,
					}
					if id, err := s.Db.InsertServiceTag(svcTag); err != nil {
						return err
					} else {
						if s.verbose {
							fmt.Println("service tag: ", tag.Title, "; id: ", id)
						}
					}
				}

				for _, image := range service.Images {
					svcImage := &request.ServicePhotoModel{
						Vendor:    service.Vendor,
						ServiceId: id,
						Url:       image.Url,
					}

					if _, err := s.Db.InsertServicePhoto(svcImage); err != nil {
						return err
					} else {
						if s.verbose {
							fmt.Println("service image: ", image.Url, "; id: ", id)
						}
					}
				}

				for _, review := range service.Reviews {
					svcReview := &request.ReviewModel{
						ServiceId: id,
						Rating:    review.Rating,
					}

					if id, err := s.Db.InsertReview(svcReview); err != nil {
						return err
					} else {
						if s.verbose {
							fmt.Println("rating: ", review.Rating, "; id: ", id)
						}
					}
				}
			}
		default:
			println("unsupported table")
		}
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
