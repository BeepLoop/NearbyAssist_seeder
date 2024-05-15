package database

import "github.com/BeepLoop/nearbyassist_seeder/request"

type Database interface {
	InitConnection() error

	InsertTag(tag *request.TagModel) (int, error)
	InsertAdmin(admin *request.AdminModel) (int, error)
	InsertUser(user *request.UserModel) (int, error)
	InsertVendor(vendor *request.VendorModel) (int, error)
	InsertService(service *request.ServiceModel) (int, error)
	InsertServiceTag(serviceTag *request.ServiceTagModel) (int, error)
	InsertReview(review *request.ReviewModel) (int, error)
	InsertServicePhoto(photo *request.ServicePhotoModel) (int, error)
}
