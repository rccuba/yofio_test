package db

import (
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"sync"
	"test_robert_yofio/internal/static"
)

// MongoConnection
type MongoConnection struct {
	*mgo.Session
	m sync.Mutex
}

var BdName = viper.GetString(static.MONGO_DATABASE)

//NewConnection
func NewConnection(info *mgo.DialInfo) (*MongoConnection, error) {

	sess, err := mgo.DialWithInfo(info)

	if err != nil || sess == nil {
		return nil, err
	}

	sess.SetMode(mgo.Monotonic, true)
	sess.SetSafe(&mgo.Safe{})

	return &MongoConnection{sess, sync.Mutex{}}, nil
}

//Create Connection
func (con *MongoConnection) CreateConnection() *mgo.Database {
	con.m.Lock()
	defer con.m.Unlock()
	return con.DB(BdName)
}

//Finding Data
func (con *MongoConnection) GetFindData(collection string, query bson.M, selector bson.M, fieldSort string, orderSort string) ([]interface{}, error) {
	c := con.CreateConnection().C(collection)
	result := make([]interface{}, 0)
	err := c.Find(query).Select(selector).Sort(orderSort + fieldSort).All(&result)
	return result, err
}

//Inserting Data
func (con *MongoConnection) InsertData(collection string, ui interface{}) (err error) {
	c := con.CreateConnection().C(collection)
	event := ui
	err = c.Insert(&event)
	return err
}
