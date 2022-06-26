package mysql

const (
	sqlGetAll = "SELECT  * FROM movie"

	sqlGetById = "SELECT * FROM movie WHERE id=?"

	sqlCreate = `
	INSERT INTO 
	movie (name, genre, year, award) 
	VALUES (?, ?, ?, ?)
	`

	sqlUpdateAward = `
	UPDATE movie 
	SET award=?
	WHERE id=?
	`

	sqlDelete = "DELETE FROM product WHERE id=?"
)
