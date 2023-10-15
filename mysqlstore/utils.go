package mysqlstore

import "gorm.io/gorm"

func Wheres(db *gorm.DB, filters ...*Filter) *gorm.DB {
	if len(filters) == 0 {
		return db
	}

	for _, f := range filters {
		db = db.Where(f.Query, f.Args...)
	}

	return db
}
