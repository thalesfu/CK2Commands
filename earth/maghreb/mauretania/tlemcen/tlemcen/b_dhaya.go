package tlemcen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎耶DhayaBarony struct {
	feud.BaseBarony
}

var BDhaya扎耶 feud.Barony = &扎耶DhayaBarony{}

func init() {
    f := BDhaya扎耶.(*扎耶DhayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dhaya",
		TitleName: "扎耶",
		TitleCode: "b_dhaya",
	}
}
