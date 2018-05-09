package dao

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"fmt"
	"strconv"
	. "../models"
	)

	type UserDAO struct {
		Server string
		Database string
}


var db *mgo.Database

const (
	COLLECTION = "users"
)

func(m *UserDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}

	db = session.DB(m.Database)

	index := mgo.Index{
		Key: []string{"id"},
		Unique: true,
	}
	err = db.C(COLLECTION).EnsureIndex(index)
	if err != nil {
  	fmt.Println(err)
	}
}


func (m *UserDAO) FindAll() ([]User, error) {
	var users []User

	err := db.C(COLLECTION).Find(bson.M{}).All(&users)
	return users, err
}

func (m *UserDAO) FindById(id string) (User, error) {
	var user User
	user_id, err := strconv.Atoi(id)
	if err != nil {
		return user, err
	}
	err = db.C(COLLECTION).Find(bson.M{"id":user_id}).One(&user)
	return user, err
}
func (m *UserDAO) Insert(user User) error {
	err := db.C(COLLECTION).Insert(&user)
	return err
}
func (m *UserDAO) Delete(id string) error{
	user_id, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	err = db.C(COLLECTION).Remove(bson.M{"id":user_id})
	return err
}
func (m UserDAO) Update(user User, id string) error {
	user_id, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	fmt.Println(user.Name)
	user.Updated_at = bson.Now()
	err = db.C(COLLECTION).Update(bson.M{"id":user_id}, &user)
	return err
}
