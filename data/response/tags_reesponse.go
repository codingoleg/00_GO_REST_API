package response

type TagsResponse struct {
	OrderId int    `json:"order_id"`
	UserId  int    `json:"user_id"`
	Name    string `json:"name"`
}
