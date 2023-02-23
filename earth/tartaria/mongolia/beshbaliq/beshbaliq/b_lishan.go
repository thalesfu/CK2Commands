package beshbaliq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 骊山LishanBarony struct {
	feud.BaseBarony
}

var BLishan骊山 feud.Barony = &骊山LishanBarony{}

func init() {
	f := BLishan骊山.(*骊山LishanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lishan",
		TitleName: "骊山",
		TitleCode: "b_lishan",
	}
}
