package soyma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索伊马SoymaBarony struct {
	feud.BaseBarony
}

var BSoyma索伊马 feud.Barony = &索伊马SoymaBarony{}

func init() {
    f := BSoyma索伊马.(*索伊马SoymaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "soyma",
		TitleName: "索伊马",
		TitleCode: "b_soyma",
	}
}
