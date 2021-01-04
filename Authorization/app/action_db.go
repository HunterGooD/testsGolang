package app

// CreateUser создает пользователя в базе данных
func (a *App) createUser(login, passwordHash string) error {
	_, err := a.db.Exec("INSERT INTO `user` (`login`, `password`) VALUES ($1, $2)", login, passwordHash)
	return err
}
