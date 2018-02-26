package repo

import (
	"log"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/jeffersonsc/speed-url/conf"
	"github.com/jeffersonsc/speed-url/model"
	"github.com/speps/go-hashids"
)

func SaveURL(url string) (myurl model.MyURL, err error) {
	myurl = model.MyURL{}
	myurl.ID = bson.NewObjectId()
	myurl.LongURL = url
	myurl.CreatedAt = time.Now()

	hd := hashids.NewData()
	h, _ := hashids.NewWithData(hd)
	now := time.Now()
	myurl.Key, _ = h.Encode([]int{int(now.Unix())})
	myurl.ShortURL = "http://localhost:3000/" + myurl.Key

	cl, err := conf.GetMongoCollection("myurls")
	if err != nil {
		return
	}

	err = cl.Insert(&myurl)
	if err != nil {
		log.Println("[storage] Erro ao salvar url. ERROR: ", err.Error())
	}

	return
}

func FindURL(key string, ip string) (url string, err error) {
	cl, err := conf.GetMongoCollection("myurls")
	if err != nil {
		return
	}

	var myurl model.MyURL
	err = cl.Find(bson.M{"key": key}).Select(bson.M{"_id": 1, "long_url": 1, "count": 1}).One(&myurl)
	if err != nil {
		return
	}
	log.Printf("%+v", myurl)
	// Save ip on request
	ct := myurl.Count + 1
	access := bson.M{"access": bson.M{"ip": ip, "created_at": time.Now()}}
	err = cl.Update(bson.M{"_id": myurl.ID}, bson.M{"$set": bson.M{"count": ct}, "$push": access})
	if err != nil {
		return
	}

	return myurl.LongURL, nil
}

func ExplaneURL(key string) (myurl model.MyURL, err error) {

	cl, err := conf.GetMongoCollection("myurls")
	if err != nil {
		return
	}

	err = cl.Find(bson.M{"key": key}).One(&myurl)
	if err != nil {
		return
	}

	return
}
