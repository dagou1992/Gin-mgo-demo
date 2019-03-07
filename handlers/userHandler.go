package handlers

import (
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "../bean"
    "../db"
)

const (
    DATABASE = "test"
    COLLECTIONNAME = "user"
)

func withCollection(f func(*mgo.Collection) error) error {
    session := db.Connect()
    defer session.Close()
    c := session.DB(DATABASE).C(COLLECTIONNAME)
    return f(c)
}

func Insert(user bean.User) (err error) {
    err = withCollection(func(c *mgo.Collection) error {
        return c.Insert(&user)
    })
    return
}

func Find(name string) (user []bean.User, err error) {
    selector := bson.M{}
    if name != "" {
        selector = bson.M{"name": name}
    }
    err = withCollection(func(c *mgo.Collection) error {
        return c.Find(selector).All(&user)
    })
    return 
}

func Update(name string, age string, phone string) (err error) {
    data := bson.M{"age": age, "phone": phone}
    selector := bson.M{"name": name}
    updateData := bson.M{"$set": data}
    err = withCollection(func(c *mgo.Collection) error {
        return c.Update(selector, updateData)
    })
    return
}

func Delete(name string) (err error) {
    selector := bson.M{"name": name}
    err = withCollection(func(c *mgo.Collection) error {
        return c.Remove(selector)
    })
    return 
}