package barkul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 天头村TiantoucunBarony struct {
	feud.BaseBarony
}

var BTiantoucun天头村 feud.Barony = &天头村TiantoucunBarony{}

func init() {
    f := BTiantoucun天头村.(*天头村TiantoucunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tiantoucun",
		TitleName: "天头村",
		TitleCode: "b_tiantoucun",
	}
}
