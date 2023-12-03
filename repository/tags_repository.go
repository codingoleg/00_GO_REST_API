package repository

import "07_restapi/model"

type TagsRepository interface {
	Save(tags model.Tags)
	FindByOrderId(tagsId int) (tags model.Tags, err error)
	FindByUserId(tagsId int) []model.Tags
}
