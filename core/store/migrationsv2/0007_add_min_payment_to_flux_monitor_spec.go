package migrationsv2

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

const (
	up7 = `
ALTER TABLE flux_monitor_specs
ADD min_payment varchar(255);
`
	down7 = `
ALTER TABLE flux_monitor_specs
DROP COLUMN min_payment;
`
)

func init() {
	Migrations = append(Migrations, &gormigrate.Migration{
		ID: "0007_add_min_payment_to_flux_monitor_spec",
		Migrate: func(db *gorm.DB) error {
			return db.Exec(up7).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.Exec(down7).Error
		},
	})
}
