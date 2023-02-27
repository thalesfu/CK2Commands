package ochrid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德巴尔DebarBarony struct {
	feud.BaseBarony
}

var BDebar德巴尔 feud.Barony = &德巴尔DebarBarony{}

func init() {
    f := BDebar德巴尔.(*德巴尔DebarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "debar",
		TitleName: "德巴尔",
		TitleCode: "b_debar",
	}
}
