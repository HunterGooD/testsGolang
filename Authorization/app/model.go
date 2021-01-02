package app

// Authorization для афторизации и регистрации
type Authorization struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

//ResponseToken ответное сообщение при регистрации или авторизации
type ResponseToken struct {
	Token string `json:"token"`
}

// ErrorResponse ответ ошибки если что то пошло не так
type ErrorResponse struct {
	Error map[string]interface{}
}

// Package ///
type Package struct {
	Head HeadPackage `json:"head"`
	Body BodyPackage `json:"body"`
}

// HeadPackage ...
type HeadPackage struct {
	Rand       string `json:"rand"`
	Title      string `json:"title"`
	Sender     string `json:"sender"`
	SessionKey string `json:"session"`
}

//BodyPackage ...
type BodyPackage struct {
	Data string `json:"data"`
}
