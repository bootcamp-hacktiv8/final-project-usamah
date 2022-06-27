package helper

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

type SocialMediaInput struct {
	Name             string `json:"name"`
	Social_media_url string `json:"social_media_url"`
}

type PhotoInput struct {
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	Photo_url string `json:"photo_url"`
}

type CommentInput struct {
	Message string `json:"message"`
}
