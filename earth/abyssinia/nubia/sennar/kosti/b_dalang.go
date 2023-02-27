package kosti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪灵DalangBarony struct {
	feud.BaseBarony
}

var BDalang迪灵 feud.Barony = &迪灵DalangBarony{}

func init() {
    f := BDalang迪灵.(*迪灵DalangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dalang",
		TitleName: "迪灵",
		TitleCode: "b_dalang",
	}
}
