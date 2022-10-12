package request

type RegisterUserOption struct {
	Username             string `json:"username" validate:"required,min=2,max=30"`
	Password             string `json:"password" validate:"required,min=6,max=20,eqfield=PasswordConfirmation"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,min=6,max=20"`
}

type UserLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CreateTagOption struct {
	Name   string `json:"name" validate:"required,min=1,max=30"`
	UserID int    `json:"user_id"`
}

type CreateArticleTagOption struct {
	ArticleID int   `json:"article_id" validate:"required"`
	TagIDs    []int `json:"tag_ids" validate:"required"`
	UserID    int   `json:"user_id"`
}
