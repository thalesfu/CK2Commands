package lopnor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 南屯NantungBarony struct {
	feud.BaseBarony
}

var BNantung南屯 feud.Barony = &南屯NantungBarony{}

func init() {
	f := BNantung南屯.(*南屯NantungBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nantung",
		TitleName: "南屯",
		TitleCode: "b_nantung",
	}
}
