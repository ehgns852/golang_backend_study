package types

type User struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

type GetUserResponse struct {
	*ApiResponse
	Users []*User `json:"result"`
}

type CreateUserRequest struct {
	Name string `json:"name" binding:"required"`
	Age  int64  `json:"age" binding:"required"`
}

func (c *CreateUserRequest) ToUser() *User {
	return &User{
		Name: c.Name,
		Age:  c.Age,
	}
}

func (c *DeleteUserRequest) ToUser() *User {
	return &User{
		Name: c.Name,
	}
}

type UpdateUserRequest struct {
	Name       string `json:"name" binding:"required"`
	UpdatedAge int64  `json:"updatedAge" binding:"required"`
}

type DeleteUserRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreateUserResponse struct {
	*ApiResponse
}

type UpdateUserResponse struct {
	*ApiResponse
}

type DeleteUserResponse struct {
	*ApiResponse
}
