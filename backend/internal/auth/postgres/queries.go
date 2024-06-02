package postgres

const (
	queryRegisterUser = "INSERT INTO users (login, email, password) VALUES ($1, $2, $3);"
)