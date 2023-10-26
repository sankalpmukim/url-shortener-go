package database

func (db *Database) GetLinks() ([]Link, error) {
	var links []Link
	err := db.conn.Select(&links, "SELECT * FROM links")
	return links, err
}

func (db *Database) GetLinkByEndpoint(endpoint string) (Link, error) {
	var link Link
	err := db.conn.Get(&link, "SELECT * FROM links WHERE endpoint=$1", endpoint)
	return link, err
}

func (db *Database) LinkExists(target string) bool {
	var link Link
	err := db.conn.Get(&link, "SELECT * FROM links WHERE target=$1", target)
	return err == nil
}

func (db *Database) UpdateLink(oldEndpoint, newEndpoint, target string) error {
	_, err := db.conn.Exec("UPDATE links SET endpoint=$1, target=$2, updatedat=NOW() WHERE endpoint=$3", newEndpoint, target, oldEndpoint)
	return err
}

func (db *Database) CreateLink(link CreateLink) error {
	_, err := db.conn.Exec("INSERT INTO links (endpoint, target, clicks, createdby, createdat, updatedat) VALUES ($1, $2, $3, $4, NOW(), NOW())", link.Endpoint, link.Target, 0, link.CreatedBy)
	return err
}

func (db *Database) IncrementClicks(endpoint string) error {
	_, err := db.conn.Exec("UPDATE links SET clicks=clicks+1, updatedat=NOW() WHERE endpoint=$1", endpoint)
	return err
}
