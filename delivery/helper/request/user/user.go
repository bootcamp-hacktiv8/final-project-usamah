package request

type FormatLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
