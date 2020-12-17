package app

type authorization struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type responseToken struct {
	Token string `json:"token"`
}
