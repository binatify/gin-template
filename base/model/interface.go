package model

import (
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// MongoDB record interface defination
type Record interface {
	IsNewRecord() bool
	IsValid() bool
	C() Collection

	id() bson.ObjectId
	setIsNewRecord(bool)
	setCreatedAt(time.Time)
	setUpdatedAt(time.Time)
}

// db Collection interface defination
type Collection interface {
	Query(query func(c *mgo.Collection))
}

func Save(r Record, migrations ...bson.M) (err error) {
	if !r.id().Valid() {
		return ErrInvalidID
	}

	if !r.IsValid() {
		return ErrInvalidArgs
	}

	r.C().Query(func(mc *mgo.Collection) {
		t := time.Now()
		if r.IsNewRecord() {
			r.setCreatedAt(t)
			r.setUpdatedAt(t)

			err = mc.Insert(r)
			if err == nil {
				r.setIsNewRecord(false)
			}

		} else {
			var migration bson.M

			if len(migrations) > 0 {
				migration = migrations[0]
			} else {
				migration = bson.M{}
			}

			migration["updated_at"] = t

			err = mc.UpdateId(r.id(), bson.M{
				"$set": migration,
			})
		}
	})
	return
}

func Destroy(r Record) (err error) {
	r.C().Query(func(mc *mgo.Collection) {
		err = mc.RemoveId(r.id())
	})
	return
}

func Find(c Collection, id string, r interface{}) (err error) {
	if !bson.IsObjectIdHex(id) {
		return ErrInvalidID
	}

	c.Query(func(mc *mgo.Collection) {
		err = mc.FindId(bson.ObjectIdHex(id)).One(r)
	})

	return
}

func FindBy(c Collection, query bson.M, r interface{}) (err error) {
	c.Query(func(mc *mgo.Collection) {
		err = mc.Find(query).One(r)
	})

	return
}

func Where(c Collection, query bson.M, total int, r interface{}, sort ...string) (err error) {
	sortStr := "-_id"
	if len(sort) > 0 {
		sortStr = sort[0]
	}

	c.Query(func(mc *mgo.Collection) {
		if total == -1 {
			err = mc.Find(query).Sort(sortStr).All(r)
		} else {
			err = mc.Find(query).Limit(total).Sort(sortStr).All(r)
		}
	})

	return
}

func Update(r Record, query bson.M, update bson.M) (err error) {
	if r.IsNewRecord() {
		err = ErrNotPersisted
		return
	}

	r.C().Query(func(c *mgo.Collection) {
		err = c.Update(query, bson.M{
			"$set": update,
		})
	})

	return
}

func AllBy(c Collection, nPerPage, pageNum int, query bson.M, r interface{}, sortStr ...string) (err error) {
	skipNum := 0
	if pageNum > 0 {
		skipNum = (pageNum - 1) * nPerPage
	}

	var sort string
	if len(sortStr) == 0 {
		sort = "-_id"
	} else {
		sort = sortStr[0]
	}

	c.Query(func(mc *mgo.Collection) {
		err = mc.Find(query).Skip(skipNum).Limit(nPerPage).Sort(sort).All(r)
	})
	return
}

func BatchInsert(c Collection, docs ...interface{}) (err error) {
	c.Query(func(mc *mgo.Collection) {
		err = mc.Insert(docs...)
	})

	return
}

func BatchDelete(c Collection, query bson.M) (err error) {
	c.Query(func(mc *mgo.Collection) {
		_, err = mc.RemoveAll(query)
	})

	return
}

func CountNum(c Collection, query ...bson.M) (num int, err error) {
	if len(query) > 0 {
		c.Query(func(mc *mgo.Collection) {
			num, err = mc.Find(query[0]).Count()
		})
	} else {
		c.Query(func(mc *mgo.Collection) {
			num, err = mc.Count()
		})
	}
	return
}
