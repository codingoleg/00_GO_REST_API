package service

import (
	"07_restapi/data/request"
	"07_restapi/data/response"
)

type TagsService interface {
	Create(tags request.CreateTagsRequest)
	FindByOrderId(tagsId int) response.TagsResponse
	FindByUserId(tagsId int) []response.TagsResponse
	FindAll() []response.TagsResponse
}
