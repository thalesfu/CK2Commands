package avranches

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴夫勒尔BarfleurBarony struct {
	feud.BaseBarony
}

var BBarfleur巴夫勒尔 feud.Barony = &巴夫勒尔BarfleurBarony{}

func init() {
	f := BBarfleur巴夫勒尔.(*巴夫勒尔BarfleurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barfleur",
		TitleName: "巴夫勒尔",
		TitleCode: "b_barfleur",
	}
}
