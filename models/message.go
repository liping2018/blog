// file ：message.go
// auth : Li ping
// date : 2018-09-24
// desc : 评论实体

package models

import (
	"time"
)

type Message struct {
	Messageid int       `orm:"column(messageid);auto"  description:"" json:"messageid"`
	Content   string    `orm:"column(content);null"  description:"" json:"content"`
	Sendtime  time.Time `orm:"column(sendtime);type(datetime);null"  description:"" json:"sendtime"`
	Userid    int       `orm:"column(userid);null"  description:"" json:"userid"`
	Articleid int       `orm:"column(articleid);null"  description:"" json:"articleid"`
}

func (t *Message) TableName() string {
	return "message"
}
