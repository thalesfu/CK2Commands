package kairwan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈富兹HaffouzBarony struct {
	feud.BaseBarony
}

var BHaffouz哈富兹 feud.Barony = &哈富兹HaffouzBarony{}

func init() {
	f := BHaffouz哈富兹.(*哈富兹HaffouzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "haffouz",
		TitleName: "哈富兹",
		TitleCode: "b_haffouz",
	}
}
