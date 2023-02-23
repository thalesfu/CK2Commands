package saur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达伊尔DairBarony struct {
	feud.BaseBarony
}

var BDair达伊尔 feud.Barony = &达伊尔DairBarony{}

func init() {
	f := BDair达伊尔.(*达伊尔DairBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dair",
		TitleName: "达伊尔",
		TitleCode: "b_dair",
	}
}
