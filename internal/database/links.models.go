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
