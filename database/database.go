package database

import (
	"github.com/kwanok/minsim-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type DB interface {
	Session() *gorm.DB
	Migrate()
}

type db struct {
	session *gorm.DB
}

func (d *db) Session() *gorm.DB {
	return d.session
}

func (d *db) Migrate() {
	models := []interface{}{
		&RedditPost{},
		&RedditComment{},
		&Minsim{},
	}

	for _, model := range models {
		err := d.session.AutoMigrate(model)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func buildDsn(c *config.Config) string {
	return "host=" + c.Database.Host +
		" user=" + c.Database.User +
		" password=" + c.Database.Password +
		" dbname=" + c.Database.Database +
		" port=" + c.Database.Port +
		" sslmode=disable"
}

func NewDB(
	c *config.Config,
) DB {
	gdb, err := gorm.Open(postgres.Open(buildDsn(c)), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &db{
		session: gdb,
	}
}
