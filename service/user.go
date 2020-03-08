package service

import (
	"github.com/xavient/crud/models"
	"log"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)




type DB struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "movies"
)

// Establish a connection to database
func (m *DB) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of users
func (m *DB) FindAll() (u []models.User, err error) {
	err = db.C(COLLECTION).Find(bson.M{}).All(&u)
	return u, err
}

// Find a user by its id
func (m *DB) FindById(id string) (u models.User, err error) {
	err = db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&u)
	return u, err
}

// Insert a user into database
func (m *DB) Insert(u models.User) (err error) {
	err = db.C(COLLECTION).Insert(&u)
	return err
}

// Delete an existing user
func (m *DB) Delete(u models.User) (err error) {
	err = db.C(COLLECTION).Remove(&u)
	return err
}

// Update an existing user
func (m *DB) Update(u models.User) (err error) {
	err = db.C(COLLECTION).UpdateId(u.ID, &u)
	return err
}
