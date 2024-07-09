package daos

import "time"

type Url struct {
	Id      string
	UserId  uint64
	Url     string
	Created time.Time
	Active  bool
}

func (q RWQuery) CreateUrl(url Url) error {
	_, err := q.db.Exec(
		"INSERT INTO url VALUES (?, ?, ?, ?, ?)",
		url.Id, url.UserId, url.Url, url.Created, url.Active)
	return err
}

func (q Query) GetUrl(id string) (string, error) {
	var url string
	row := q.db.QueryRow("SELECT url FROM url WHERE id == ?", id)
	err := row.Scan(&url)
	return url, err
}

func (q Query) GetUrlData(id string) (*Url, *User, error) {
	url := Url{}
	user := User{}
	row := q.db.QueryRow(`
		SELECT url.id, url.url, url.created, url.active, url.user_id, mu.id, mu.name, mu.created, mu.active
		FROM url JOIN m_user mu on url.user_id = mu.id WHERE url.id == ?`, id)
	err := row.Scan(&url.Id, &url.Url, &url.Created, &url.Active,
		&url.UserId, &user.Id, &user.Name, &user.Created, &user.Active)
	if err != nil {
		return nil, nil, err
	}
	return &url, &user, nil
}

func (q RWQuery) UpdateUrlAccessibility(id string, active bool) error {
	_, err := q.db.Exec("UPDATE url SET active=? WHERE id == ?", active, id)
	return err
}
