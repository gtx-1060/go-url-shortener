package daos

import "time"

type User struct {
	Id      uint64
	Name    string
	Created time.Time
	Active  bool
}

func (q RWQuery) CreateUser(user User) (uint64, error) {
	var id uint64
	row := q.db.QueryRow(
		"INSERT INTO m_user VALUES (NULL, ?, ?, ?) RETURNING id",
		user.Name, user.Created, user.Active)
	err := row.Scan(&id)
	return id, err
}

func (q Query) GetUser(name string) (*User, error) {
	user := &User{Name: name}
	row := q.db.QueryRow("SELECT id, active, created FROM m_user WHERE name == ?", name)
	if err := row.Scan(&user.Id, &user.Active, &user.Created); err != nil {
		return nil, err
	}
	return user, nil
}

func (q RWQuery) UpdateUserAccessibility(id uint64, active bool) error {
	_, err := q.db.Exec("UPDATE m_user SET active=? WHERE id == ?", active, id)
	return err
}
