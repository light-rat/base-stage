package response

type GetUserListResponse struct {
	List  []UserListItem `json:"list"`
	Count int64          `json:"total"`
}

type UserListItem struct {
	Name  string `json:"name"`
	Home  string `json:"home"`
	Email string `json:"email"`
	Sex   int    `json:"sex"`
}
