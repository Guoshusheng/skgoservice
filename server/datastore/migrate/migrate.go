package migrate

import (
	"github.com/BurntSushi/migration"
)

// Setup is the database migration function that
// will setup the initial SQL database structure.
func Setup(tx migration.LimitedTx) error {
	var stmts = []string{
		testskTable,
	}
	for _, stmt := range stmts {
		_, err := tx.Exec(transform(stmt))
		if err != nil {
			return err
		}
	}
	return nil
}

//// Migrate_20142110 is a database migration on Oct-10 2014.
//func Migrate_20142110(tx migration.LimitedTx) error {
//	var stmts = []string{
//		commitRepoIndex, // index the commit table repo_id column
//		repoTokenColumn, // add the repo token column
//		repoTokenUpdate, // update the repo token column to empty string
//	}
//	for _, stmt := range stmts {
//		_, err := tx.Exec(transform(stmt))
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}

//// Migrate_20142110 is a database migration on Oct-10 2014.
//func Migrate_20152701(tx migration.LimitedTx) error {
//	var stmts = []string{
//		addUserTokenExpires, // index the commit table repo_id column
//	}
//	for _, stmt := range stmts {
//		_, err := tx.Exec(transform(stmt))
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}

var testskTable = `
CREATE TABLE IF NOT EXISTS testsks (
	 testsk_id         INTEGER PRIMARY KEY AUTOINCREMENT
	,testsk_time       VARCHAR(255)
	,testsk_add 	   VARCHAR(255)
	,testsk_name 	   VARCHAR(255)
	,testsk_count	   INTEGER
	,UNIQUE(testsk_name)
);
`
