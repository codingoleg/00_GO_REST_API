package service

import (
	"07_restapi/data/request"
	"07_restapi/data/response"
	"07_restapi/helper"
	"07_restapi/model"
	"07_restapi/repository"

	"github.com/go-playground/validator/v10"
)

type TagsServiceImpl struct {
	TagsRepository repository.TagsRepository
	Validate       *validator.Validate
}

func NewTagsServiceImpl(tagRepository repository.TagsRepository, validate *validator.Validate) TagsService {
	return &TagsServiceImpl{
		TagsRepository: tagRepository,
		Validate:       validate,
	}
}

func (t *TagsServiceImpl) Create(tags request.CreateTagsRequest) {
	err := t.Validate.Struct(tags)
	helper.ErrorPanic(err)
	tagModel := model.Tags{
		OrderId: tags.OrderId,
		UserId:  tags.UserId,
		Name:    tags.Name,
	}
	t.TagsRepository.Save(tagModel)
}

func (t *TagsServiceImpl) FindByOrderId(tagsId int) response.TagsResponse {
	tagData, err := t.TagsRepository.FindByOrderId(tagsId)
	helper.ErrorPanic(err)

	tagResponse := response.TagsResponse{
		OrderId: tagData.OrderId,
		UserId:  tagData.UserId,
		Name:    tagData.Name,
	}
	return tagResponse
}

func (t *TagsServiceImpl) FindByUserId(tagsId int) []response.TagsResponse {
	result := t.TagsRepository.FindByUserId(tagsId)

	var tags []response.TagsResponse
	for _, value := range result {
		tag := response.TagsResponse{
			OrderId: value.OrderId,
			UserId:  value.UserId,
			Name:    value.Name,
		}
		tags = append(tags, tag)
	}

	return tags
}

func (t *TagsServiceImpl) FindAll() []response.TagsResponse {
	result := t.TagsRepository.FindAll()

	var tags []response.TagsResponse
	for _, value := range result {
		tag := response.TagsResponse{
			OrderId: value.OrderId,
			UserId:  value.UserId,
			Name:    value.Name,
		}
		tags = append(tags, tag)
	}

	return tags
}
