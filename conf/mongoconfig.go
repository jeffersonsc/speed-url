package conf

import (
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
)

var mgoSession *mgo.Session
var mgoAuth *mgo.DialInfo
var mongoDbURI string
var mongoDbName string
var mongoDbUser string
var mongoDbPass string

/*
LoadMongoConfig Init MongoDB Connection
*/
func LoadMongoConfig() {
	mongoDbURI = Cfg.Section("").Key("mongo_uri").Value()
	mongoDbName = Cfg.Section("").Key("mongo_db").Value()
	mongoDbUser = Cfg.Section("").Key("mongo_user").Value()
	mongoDbPass = Cfg.Section("").Key("mongo_pass").Value()

	mgoAuth = new(mgo.DialInfo)
	mgoAuth.Addrs = []string{mongoDbURI}
	mgoAuth.Timeout = time.Second * 60
	mgoAuth.Database = mongoDbName
	mgoAuth.Username = mongoDbUser
	mgoAuth.Password = mongoDbPass
}

/*
GetMongoSession gets connection to Mongo repo
*/
func GetMongoSession() (*mgo.Session, error) {
	if mgoSession != nil {
		mgoSession.Refresh()
		return mgoSession.Copy(), nil
	}

	LoadMongoConfig()
	mgoSession, erro := mgo.DialWithInfo(mgoAuth)
	if erro != nil {
		log.Printf("[GetMongoSession] Erro ao tentar abrir a sessao com o Mongo: [%s]\n", erro.Error())
		return nil, erro
	}
	return mgoSession.Copy(), erro
}

/*
GetMongoCollection gets a data collection
*/
func GetMongoCollection(collectionName string) (*mgo.Collection, error) {
	mgoSession, erro := GetMongoSession()
	if erro != nil {
		log.Printf("[GetCollection] Erro ao conectar ao Mongo: [%s]\n", erro.Error())
		return nil, erro
	}

	collection := mgoSession.DB(mongoDbName).C(collectionName)
	return collection, nil
}
