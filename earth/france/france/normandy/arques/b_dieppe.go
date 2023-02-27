package arques

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪耶普DieppeBarony struct {
	feud.BaseBarony
}

var BDieppe迪耶普 feud.Barony = &迪耶普DieppeBarony{}

func init() {
    f := BDieppe迪耶普.(*迪耶普DieppeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dieppe",
		TitleName: "迪耶普",
		TitleCode: "b_dieppe",
	}
}
