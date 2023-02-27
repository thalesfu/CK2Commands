package kanj_rustaq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪扎DizaBarony struct {
	feud.BaseBarony
}

var BDiza迪扎 feud.Barony = &迪扎DizaBarony{}

func init() {
    f := BDiza迪扎.(*迪扎DizaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "diza",
		TitleName: "迪扎",
		TitleCode: "b_diza",
	}
}
