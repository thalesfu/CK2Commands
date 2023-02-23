package luristan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德巴尔DezbarBarony struct {
	feud.BaseBarony
}

var BDezbar德巴尔 feud.Barony = &德巴尔DezbarBarony{}

func init() {
	f := BDezbar德巴尔.(*德巴尔DezbarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dezbar",
		TitleName: "德巴尔",
		TitleCode: "b_dezbar",
	}
}
