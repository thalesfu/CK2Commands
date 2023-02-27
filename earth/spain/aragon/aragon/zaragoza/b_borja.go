package zaragoza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博尔哈BorjaBarony struct {
	feud.BaseBarony
}

var BBorja博尔哈 feud.Barony = &博尔哈BorjaBarony{}

func init() {
    f := BBorja博尔哈.(*博尔哈BorjaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "borja",
		TitleName: "博尔哈",
		TitleCode: "b_borja",
	}
}
