package repository

import (
	"07_restapi/helper"
	"07_restapi/model"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type TagsRepositoryImpl struct {
	Db *gorm.DB
}

func NewTagsRepositoryImpl(Db *gorm.DB) TagsRepository {
	return &TagsRepositoryImpl{Db: Db}
}

func (t *TagsRepositoryImpl) Save(tags model.Tags) {
	result := t.Db.Create(&tags)
	helper.ErrorPanic(result.Error)
}

func (t *TagsRepositoryImpl) FindByOrderId(tagsId int) (tags model.Tags, err error) {
	var tag model.Tags
	result := t.Db.Find(&tag, tagsId)
	if result != nil {
		return tag, nil
	} else {
		return tag, errors.New("order is not found")
	}
}

func (t *TagsRepositoryImpl) FindByUserId(tagsId int) []model.Tags {
	var tags []model.Tags
	result := t.Db.Where("user_id = ?", tagsId).Find(&tags)
	for _, user := range tags {
		fmt.Println(user)
	}

	helper.ErrorPanic(result.Error)
	return tags
}

func (t *TagsRepositoryImpl) FindAll() []model.Tags {
	var tags []model.Tags
	result := t.Db.Find(&tags)
	helper.ErrorPanic(result.Error)
	return tags
}
