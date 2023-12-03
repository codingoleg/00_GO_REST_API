package request

type CreateTagsRequest struct {
	OrderId int    `json:"order_id"`
	UserId  int    `json:"user_id"`
	Name    string `validate:"required,min=1,max=200" json:"name"`
}
