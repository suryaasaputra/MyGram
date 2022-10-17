package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name"  form:"name" valid:"required~Name is required"`
	SocialMediaUrl string `gorm:"not null" json:"social_media_url"  form:"social_media_url" valid:"required~Social media url is required"`
	UserID         int    `json:"user_id"`
	User           *User
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(s)
	if err != nil {
		return err
	}

	return nil
}
func (s *SocialMedia) BeforeUpdate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(s)
	if err != nil {
		return err
	}

	return nil
}
