package db

func GetUsers() ([]string, error) {
	var users []string
	rows := conn.QueryRow("select name from users")
	err := rows.Scan(&users)
	return users, err
}
