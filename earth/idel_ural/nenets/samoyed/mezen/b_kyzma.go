package mezen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克兹马KyzmaBarony struct {
	feud.BaseBarony
}

var BKyzma克兹马 feud.Barony = &克兹马KyzmaBarony{}

func init() {
    f := BKyzma克兹马.(*克兹马KyzmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kyzma",
		TitleName: "克兹马",
		TitleCode: "b_kyzma",
	}
}
