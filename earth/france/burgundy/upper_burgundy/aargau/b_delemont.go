package aargau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德莱蒙DelemontBarony struct {
	feud.BaseBarony
}

var BDelemont德莱蒙 feud.Barony = &德莱蒙DelemontBarony{}

func init() {
    f := BDelemont德莱蒙.(*德莱蒙DelemontBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "delemont",
		TitleName: "德莱蒙",
		TitleCode: "b_delemont",
	}
}
