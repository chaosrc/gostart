package main

import (
	"database/sql"
	"fmt"
)

func listTrack(db sql.DB, artist string, minYear, maxYear int) {
	result, err := db.Exec("SELECT * FROM track WHERE artist = ? and year > ? and year < ? ",
		artist, minYear, maxYear)

	fmt.Println(result, err)

}

func sqlQuote(x interface{}) string {
	if x == nil {
		return "NULL"
	} else if _, ok := x.(int); ok {
		return fmt.Sprintf("%d", x)
	} else if _, ok := x.(string); ok {
		return fmt.Sprintf("%s", x)
	}

	return ""
}

func sw(x interface{}) string {
	switch x := x.(type) {
	case int, uint:
		return fmt.Sprintf("%d", x)
	case string:
		return fmt.Sprintf("%s", x)
	case bool:
		//...
	default:

	}
	return ""
}
