// GORM simplifies database interactions by mapping Go structs to database tables.

package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string
	Body  string
}
