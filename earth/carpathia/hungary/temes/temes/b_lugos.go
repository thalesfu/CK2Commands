package temes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢戈斯LugosBarony struct {
	feud.BaseBarony
}

var BLugos卢戈斯 feud.Barony = &卢戈斯LugosBarony{}

func init() {
	f := BLugos卢戈斯.(*卢戈斯LugosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lugos",
		TitleName: "卢戈斯",
		TitleCode: "b_lugos",
	}
}
