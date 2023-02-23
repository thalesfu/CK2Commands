package medelike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅尔克MelkBarony struct {
	feud.BaseBarony
}

var BMelk梅尔克 feud.Barony = &梅尔克MelkBarony{}

func init() {
	f := BMelk梅尔克.(*梅尔克MelkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "melk",
		TitleName: "梅尔克",
		TitleCode: "b_melk",
	}
}
