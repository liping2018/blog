// file ：images.go
// auth : Li ping
// date : 2018-09-24
// desc : 相册实体

package models

type Images struct {
	Imagesid  int    `orm:"column(imagesid);auto"  description:"" json:"imagesid"`
	Title     string `orm:"column(title);null"  description:"" json:"title"`
	Imagepath string `orm:"column(imagepath);null"  description:"" json:"imagepath"`
	Content   string `orm:"column(content);null"  description:"" json:"content"`
	Userid    int    `orm:"column(userid);null"  description:"" json:"userid"`
}

func (t *Images) TableName() string {
	return "images"
}
