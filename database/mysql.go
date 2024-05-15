package database

import (
	"context"
	"fmt"
	"time"

	"github.com/BeepLoop/nearbyassist_seeder/config"
	"github.com/BeepLoop/nearbyassist_seeder/request"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Mysql struct {
	conf *config.Config
	Conn *sqlx.DB
}

func NewMysqlDatabase(conf *config.Config) *Mysql {
	return &Mysql{
		conf: conf,
	}
}

func (m *Mysql) InitConnection() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", m.conf.DB_User, m.conf.DB_Pass, m.conf.DB_Host, m.conf.DB_Port, m.conf.DB_Name)

	if conn, err := sqlx.Connect("mysql", dsn); err != nil {
		return err
	} else {
		m.Conn = conn
	}

	return nil
}

func (m *Mysql) InsertTag(tag *request.TagModel) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "INSERT INTO Tag (title) VALUES (:title)"
	res, err := m.Conn.NamedExecContext(ctx, query, tag)
	if err != nil {
		return 0, err
	}

	insertId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	if ctx.Err() == context.DeadlineExceeded {
		return 0, context.DeadlineExceeded
	}

	return int(insertId), nil
}

func (m *Mysql) InsertAdmin(admin *request.AdminModel) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "INSERT INTO Admin (email, password, role) VALUES (:email, :password, :role)"
	res, err := m.Conn.NamedExecContext(ctx, query, admin)
	if err != nil {
		return 0, err
	}

	insertId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	if ctx.Err() == context.DeadlineExceeded {
		return 0, context.DeadlineExceeded
	}

	return int(insertId), nil
}

func (m *Mysql) InsertUser(user *request.UserModel) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "INSERT INTO User (name, email, imageUrl) VALUES (:name, :email, :imageUrl)"
	res, err := m.Conn.NamedExecContext(ctx, query, user)
	if err != nil {
		return 0, err
	}

	insertId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	if ctx.Err() == context.DeadlineExceeded {
		return 0, context.DeadlineExceeded
	}

	return int(insertId), nil
}

func (m *Mysql) InsertVendor(vendor *request.VendorModel) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
        INSERT INTO 
            Vendor (vendorId, job)
        VALUES
            (
                (SELECT id FROM User WHERE name = :name),
                :job
            )
    `
	res, err := m.Conn.NamedExecContext(ctx, query, vendor)
	if err != nil {
		return 0, err
	}

	insertId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	if ctx.Err() == context.DeadlineExceeded {
		return 0, context.DeadlineExceeded
	}

	return int(insertId), nil
}

func (m *Mysql) InsertService(service *request.ServiceModel) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
        INSERT INTO
            Service
                (vendorId, description, rate, latitude, longitude)
        VALUES 
            (
                :vendorId,
                :description,
                :rate,
                :latitude,
                :longitude
            )
    `
	res, err := m.Conn.NamedExecContext(ctx, query, service)
	if err != nil {
		return 0, err
	}

	insertId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	if ctx.Err() == context.DeadlineExceeded {
		return 0, context.DeadlineExceeded
	}

	return int(insertId), nil
}

func (m *Mysql) InsertServiceTag(serviceTag *request.ServiceTagModel) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
        INSERT INTO 
            Service_Tag (serviceId, tagId)
        VALUES
            (
                :serviceId,
                (SELECT id FROM Tag WHERE title = :tagTitle)
            )
    `
	res, err := m.Conn.NamedExecContext(ctx, query, serviceTag)
	if err != nil {
		return 0, err
	}

	insertId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	if ctx.Err() == context.DeadlineExceeded {
		return 0, context.DeadlineExceeded
	}

	return int(insertId), nil
}

func (m *Mysql) InsertReview(review *request.ReviewModel) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "INSERT INTO Review (serviceId, rating) VALUES (:serviceId, :rating)"
	res, err := m.Conn.NamedExecContext(ctx, query, review)
	if err != nil {
		return 0, err
	}

	insertId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	if ctx.Err() == context.DeadlineExceeded {
		return 0, context.DeadlineExceeded
	}

	return int(insertId), nil
}

func (m *Mysql) InsertServicePhoto(photo *request.ServicePhotoModel) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "INSERT INTO ServicePhoto (vendorId, serviceId, url) VALUES (:vendorId, :serviceId, :url)"
	res, err := m.Conn.NamedExecContext(ctx, query, photo)
	if err != nil {
		return 0, err
	}

	insertId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	if ctx.Err() == context.DeadlineExceeded {
		return 0, context.DeadlineExceeded
	}

	return int(insertId), nil
}
