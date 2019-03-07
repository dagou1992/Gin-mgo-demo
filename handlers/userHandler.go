package handlers

import (
    "gopkg.in/mgo.v2/bson"
    "../bean"
    "../db"
)

const (
    COLLECTIONNAME = "user"
)

func Insert(user bean.User) (err error) {
    db := db.Connect()
    err = db.C(COLLECTIONNAME).Insert(&user)
    return
}

func Find(name string) (user []bean.User, err error) {
    db := db.Connect()
    client := db.C(COLLECTIONNAME)
    if name == "" {
        err = client.Find(bson.M{}).All(&user)
    } else {
        var u bean.User
        err = client.Find(bson.M{"name": name}).One(&u)
        user = append(user, u)
    }
    return 
}

func Update(name string, age string, phone string) (err error) {
    db := db.Connect()
    data := bson.M{"age": age, "phone": phone}
    selector := bson.M{"name": name}
    updateData := bson.M{"$set": data}
    err = db.C(COLLECTIONNAME).Update(selector, updateData)
    return
}

func Delete(name string) (err error) {
    db := db.Connect()
    selector := bson.M{"name": name}
    err = db.C(COLLECTIONNAME).Remove(selector)
    return 
}