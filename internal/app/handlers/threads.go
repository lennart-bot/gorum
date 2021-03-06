package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/NathanNr/gorum/internal/pkg/config"
	"github.com/NathanNr/gorum/internal/pkg/db"
)

// Threads handler
func Threads(data HandlerData) interface{} {
	var err error

	// get board id and check if provided
	boardID := data.Request.GetInt("boardID")
	if boardID == 0 {
		// no board id provided
		return errors.New("400")
	}

	// query db
	var rows *sql.Rows
	rows, err = db.DB.Query(`SELECT threads.id, threads.threadname, threads.author, threads.created, users.username
							FROM threads INNER JOIN users ON threads.author = users.id WHERE threads.board = $1;`, boardID)

	// defer close and check for error
	defer rows.Close()
	if err != nil {
		// return error
		return err
	}

	// threads list to write
	threads := map[string]interface{}{}

	// loop through threads
	for rows.Next() {
		// scan
		var id, author int
		var created int64
		var name, authorName string
		err = rows.Scan(&id, &name, &author, &created, &authorName)
		if err != nil {
			// return error
			return err
		}

		// thread map to append
		thread := map[string]interface{}{}
		thread["id"] = id
		thread["name"] = name
		thread["created"] = created
		thread["author"] = author
		thread["authorName"] = authorName

		// add avatar
		avatarPath := fmt.Sprintf("%v/%v.png", config.Get("data", "avatar"), author)
		_, err = os.Open(avatarPath)
		if os.IsNotExist(err) {
			thread["authorAvatar"] = fmt.Sprintf("%s/default", config.Get("data", "avatar"))
		} else {
			thread["authorAvatar"] = avatarPath
		}

		// append thread to threads map
		threads[strconv.Itoa(id)] = thread
	}

	// write map
	return threads
}
