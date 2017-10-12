// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package sqlite3

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

// Sqlite3DatastoreType defines all of the properties and information associated
// with connecting and talking to the database.
type Sqlite3DatastoreType struct {
	Filename string
	DB       *sql.DB
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

// New - This function will return a sqlite3.Sqlite3DatastoreType
func New(filename string) Sqlite3DatastoreType {
	var ds Sqlite3DatastoreType
	ds.Filename = filename
	return ds
}

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

// Connect - This method is called from the libstix2/datastore package and is
// used to connect to an sqlite3 database
func (ds *Sqlite3DatastoreType) Connect() error {
	var err error

	if ds.Filename == "" {
		return fmt.Errorf("A valid filename is required for connecting to the sqlite3 datastore")
	}

	err = ds.verifyFileExists()
	if err != nil {
		return err
	}

	log.Println("Connecting to sqlite3 datastore at filename", ds.Filename)

	db, sqlerr := sql.Open("sqlite3", ds.Filename)
	if sqlerr != nil {
		return fmt.Errorf("Unable to open file %s due to error: %v", ds.Filename, sqlerr)
	}
	ds.DB = db

	return nil
}

// Close - This method will close the database connection
func (ds *Sqlite3DatastoreType) Close() error {
	err := ds.DB.Close()
	if err != nil {
		return err
	}
	return nil
}

// CreateTables - This method will create all of the tables needed to store
// STIX content in the database.
func (ds *Sqlite3DatastoreType) CreateTables() {
	ds.createTable("sdo_attack_pattern", ds.attackPatternProperties())
}

// ----------------------------------------------------------------------
// Private Methods
// ----------------------------------------------------------------------

// verifyFileExists - This method will check to make sure the file is found on the filesystem
func (ds *Sqlite3DatastoreType) verifyFileExists() error {
	if _, err := os.Stat(ds.Filename); os.IsNotExist(err) {
		return fmt.Errorf("ERROR: The sqlite3 database cannot be opened due to error: %v", err)
	}
	return nil
}

// commonProperties - This method will return the the common properties
func (ds *Sqlite3DatastoreType) commonProperties() string {
	return `
	"aid" INTEGER PRIMARY KEY,
 	"stix_spec_version" TEXT NOT NULL,
 	"taxii_date_added" TEXT NOT NULL,
 	"type" TEXT NOT NULL,
 	"id" TEXT NOT NULL,
 	"created_by_ref" TEXT,
 	"created" TEXT NOT NULL,
 	"modified" TEXT NOT NULL,
 	"revoked" integer(1,0) DEFAULT 0,
 	"confidence" integer(3,0),
 	"lang" text,`
}

// attackPatternProperties  - This method will return the properties for attack patterns
func (ds *Sqlite3DatastoreType) attackPatternProperties() string {
	return ds.commonProperties() + `
	"name" text NOT NULL,
	"description" text
	`
}

// createAttackPatternTable - This method will create the actual table
func (ds *Sqlite3DatastoreType) createTable(name, properties string) {
	var stmt = `CREATE TABLE IF NOT EXISTS "` + name + `" (` + properties + `)`
	_, err := ds.DB.Exec(stmt)

	if err != nil {
		log.Println("ERROR: The", name, "table could not be created")
	}
}
