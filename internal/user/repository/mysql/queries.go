package mysql

const (
	sqlGetAll = "SELECT  * FROM user"

	sqlGetById = "SELECT * FROM user WHERE id=?"

	sqlCreate = `
	INSERT INTO 
	user (name, age, movie_genre) 
	VALUES (?, ?, ?)
	`

	sqlUpdateAge = `
	UPDATE user 
	SET age=?
	WHERE id=?
	`

	sqlDelete = "DELETE FROM user WHERE id=?"
)
