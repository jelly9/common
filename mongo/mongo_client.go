package mongo

import (
	"github.com/globalsign/mgo"
)

type Options struct {
	Url 		string
	Collection	string
	Database 	string
}

type Session struct {
	session		*mgo.Session
	collection	string
	database 	string
}

func NewSession(opt *Options) *Session {
	s, err := mgo.Dial(opt.Url)
	if err != nil {
		panic("Mgo init failed.")
	}
	return &Session{
		session: s,
		collection:	opt.Collection,
		database: opt.Database,
	}
}

func (s *Session) Session(name ...string) *mgo.Session {
	return s.session
}

func (s *Session) DB(name ...string) *mgo.Database {
	db := s.database
	if len(name) != 0 {
		db = name[0]
	}
	return s.session.DB(db)
}

func (s *Session) Collection(name ...string) *mgo.Collection {
	c := s.collection
	if len(name) != 0 {
		c = name[0]
	}
	return s.session.DB(s.database).C(c)
}

func (s *Session) Insert(docs ...interface{}) error {
	return s.session.DB(s.database).C(s.collection).Insert(docs...)
}

// Insert/Update/All/Find/One/Collection