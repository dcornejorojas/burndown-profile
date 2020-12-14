package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Profile struct {
	IDProfile uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Dni       string    `gorm:"size:100;not null" json:"dni"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Profile   string    `gorm:"size:100;not null" json:"Profile"`
	LastName  string    `gorm:"size:100;not null" json:"lastName"`
	Avatar    string    `gorm:"size:200;not null" json:"avatar"`
	Type      string    `gorm:"size:100;not null" json:"type"`
	Token     string    `gorm:"size:200;not null" json:"token"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
}

type AllProfile []Profile

func (p *Profile) Prepare() {
	p.IDProfile = 0
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.LastName = html.EscapeString(strings.TrimSpace(p.LastName))
	p.Type = html.EscapeString(strings.TrimSpace(p.LastName))
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Profile) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if p.Name == "" {
			return errors.New("Required Name")
		}
		if p.LastName == "" {
			return errors.New("Required LastName")
		}
		if p.Type == "" {
			return errors.New("Required Type of Profile")
		}

		return nil
	default:
		if p.Name == "" {
			return errors.New("Required Name")
		}
		if p.LastName == "" {
			return errors.New("Required LastName")
		}
		if p.Type == "" {
			return errors.New("Required Type of Profile")
		}
		return nil
	}
}

func (p *Profile) SaveProfile(db *gorm.DB) (*Profile, error) {
	var err error
	err = db.Debug().Create(&p).Error
	if err != nil {
		return &Profile{}, err
	}
	return p, nil
}

func (p *Profile) FindAllProfiles(db *gorm.DB) (*[]Profile, error) {
	var err error
	profiles := []Profile{}
	err = db.Debug().Model(&Profile{}).Limit(100).Find(&profiles).Error
	if err != nil {
		return &[]Profile{}, err
	}
	return &profiles, err
}

func (p *Profile) FindProfileByID(db *gorm.DB, uid uint32) (*Profile, error) {
	var err error
	err = db.Debug().Model(Profile{}).Where("id = ?", uid).Take(&p).Error
	if err != nil {
		return &Profile{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Profile{}, errors.New("Profile Not Found")
	}
	return p, err
}
