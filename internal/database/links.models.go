package database

func (db *Database) GetLinks() ([]Link, error) {
	var links []Link
	err := db.conn.Select(&links, "SELECT * FROM links")
	return links, err
}
