// file ：articel.go
// auth : Li ping
// date : 2018-09-24
// desc : 文章实体

package models

import (
	"time"
)

type Articel struct {
	Articleid int       `orm:"column(articleid);auto"  description:"" json:"articleid"`
	Userid    int       `orm:"column(userid);null"  description:"" json:"userid"`
	Title     string    `orm:"column(title);null"  description:"" json:"title"`
	Content   string    `orm:"column(content);null"  description:"" json:"content"`
	Jointime  time.Time `orm:"column(jointime);type(datetime);null"  description:"" json:"jointime"`
	Typeid    int       `orm:"column(typeid);null"  description:"" json:"typeid"`
}

func (t *Articel) TableName() string {
	return "articel"
}
