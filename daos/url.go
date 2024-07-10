package daos

import "time"

type Url struct {
	Id         string
	UserId     uint64
	Url        string
	Created    time.Time
	Expiration time.Time
}

func (q RWQuery) CreateUrl(url Url) error {
	_, err := q.db.Exec(
		"INSERT INTO url VALUES (?, ?, ?, ?, ?)",
		url.Id, url.UserId, url.Url, url.Created, url.Expiration)
	return err
}

func (q Query) GetUrl(id string) (string, error) {
	var url string
	row := q.db.QueryRow("SELECT url FROM url WHERE id == ?", id)
	err := row.Scan(&url)
	return url, err
}

func (q Query) GetActiveUrl(id string) (string, error) {
	var url string
	row := q.db.QueryRow("SELECT url FROM url WHERE id == ? AND expiration < datetime('now')", id)
	err := row.Scan(&url)
	return url, err
}

func (q Query) GetUrlData(id string) (*Url, *User, error) {
	url := Url{}
	user := User{}
	row := q.db.QueryRow(`
		SELECT url.id, url.url, url.created, url.expiration, url.user_id, mu.id, mu.name, mu.created, mu.active
		FROM url JOIN m_user mu on url.user_id = mu.id WHERE url.id == ?`, id)
	err := row.Scan(&url.Id, &url.Url, &url.Created, &url.Expiration,
		&url.UserId, &user.Id, &user.Name, &user.Created, &user.Active)
	if err != nil {
		return nil, nil, err
	}
	return &url, &user, nil
}

func (q RWQuery) UpdateUrlExpiration(id string, exp time.Time) error {
	_, err := q.db.Exec("UPDATE url SET expiration=? WHERE id == ?", exp, id)
	return err
}

func (q RWQuery) RemoveAllExpiredUrls(exp time.Time) error {
	_, err := q.db.Exec("DELETE FROM url WHERE expiration > datetime('now')", exp)
	return err
}
