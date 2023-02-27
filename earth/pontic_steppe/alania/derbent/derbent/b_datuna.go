package derbent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达图那DatunaBarony struct {
	feud.BaseBarony
}

var BDatuna达图那 feud.Barony = &达图那DatunaBarony{}

func init() {
    f := BDatuna达图那.(*达图那DatunaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "datuna",
		TitleName: "达图那",
		TitleCode: "b_datuna",
	}
}
