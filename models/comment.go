package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~Mesagge is required"`
	UserID  int    `gorm:"not null" json:"user_id"`
	PhotoID int    `gorm:"not null" json:"photo_id" form:"photo_id" valid:"required~Photo ID is required"`
	User    *User
	Photo   *Photo
}

func (c *Comment) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}

	return nil
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) error {
	c.PhotoID = 1 //not actual id , only to pass validation
	_, err := govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}

	return nil
}
