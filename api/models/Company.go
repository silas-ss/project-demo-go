package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Company struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:100;not null;unique" json:"name"`
	Token     string    `gorm:"size:100;not null;unique" json:"token"`
	Callback  string    `gorm:"size:300;not null" json:"callback"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (c *Company) Prepare() {
	c.ID = 0
	c.Name = html.EscapeString(strings.TrimSpace(c.Name))
	c.Token = uuid.NewV4().String()
	c.Callback = html.EscapeString(strings.TrimSpace(c.Callback))
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
}

func (c *Company) SaveCompany(db *gorm.DB) (*Company, error) {
	var err error
	err = db.Debug().Table("company").Model(&Company{}).Create(&c).Error
	if err != nil {
		return &Company{}, err
	}
	return c, nil
}

func (c *Company) FindAllCompanies(db *gorm.DB) (*[]Company, error) {
	var err error
	companies := []Company{}
	err = db.Debug().Table("company").Model(&Company{}).Limit(100).Find(&companies).Error
	if err != nil {
		return &[]Company{}, err
	}

	return &companies, nil
}

func (c *Company) FindCompanyByID(db *gorm.DB, companyID uint64) (*Company, error) {
	var err error
	err = db.Debug().Table("company").Model(&Company{}).Where("id = ?", companyID).Take(&c).Error
	if err != nil {
		return &Company{}, err
	}
	return c, nil
}

func (c *Company) UpdateCompany(db *gorm.DB) (*Company, error) {
	var err error

	err = db.Debug().Table("company").Model(&Company{}).Where("id = ?", c.ID).Updates(Company{Callback: c.Callback, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Company{}, err
	}
	return c, nil
}

func (c *Company) DeleteCompany(db *gorm.DB, companyId uint64) (int64, error) {
	db = db.Debug().Table("company").Model(&Company{}).Where("id = ?", companyId).Take(&Company{}).Delete(&Company{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Company not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
