package chortitza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勒塔瓦LtavaBarony struct {
	feud.BaseBarony
}

var BLtava勒塔瓦 feud.Barony = &勒塔瓦LtavaBarony{}

func init() {
    f := BLtava勒塔瓦.(*勒塔瓦LtavaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ltava",
		TitleName: "勒塔瓦",
		TitleCode: "b_ltava",
	}
}
