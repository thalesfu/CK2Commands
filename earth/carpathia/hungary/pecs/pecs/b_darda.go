package pecs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达尔达DardaBarony struct {
	feud.BaseBarony
}

var BDarda达尔达 feud.Barony = &达尔达DardaBarony{}

func init() {
	f := BDarda达尔达.(*达尔达DardaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "darda",
		TitleName: "达尔达",
		TitleCode: "b_darda",
	}
}
