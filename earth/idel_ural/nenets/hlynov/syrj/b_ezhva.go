package syrj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃日瓦EzhvaBarony struct {
	feud.BaseBarony
}

var BEzhva埃日瓦 feud.Barony = &埃日瓦EzhvaBarony{}

func init() {
    f := BEzhva埃日瓦.(*埃日瓦EzhvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ezhva",
		TitleName: "埃日瓦",
		TitleCode: "b_ezhva",
	}
}
