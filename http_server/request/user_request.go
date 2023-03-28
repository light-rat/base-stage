package request

type AddUserRequest struct {
	Name  string `json:"name"`
	Home  string `json:"home"`
	Email string `json:"email"`
	Sex   int    `json:"sex"`
}

type GetUserListRequest struct {
	AddUserRequest
}
