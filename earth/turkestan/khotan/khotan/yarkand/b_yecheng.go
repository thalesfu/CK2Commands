package yarkand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 叶城YechengBarony struct {
	feud.BaseBarony
}

var BYecheng叶城 feud.Barony = &叶城YechengBarony{}

func init() {
    f := BYecheng叶城.(*叶城YechengBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yecheng",
		TitleName: "叶城",
		TitleCode: "b_yecheng",
	}
}
