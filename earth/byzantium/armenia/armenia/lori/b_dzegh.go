package lori

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泽格DzeghBarony struct {
	feud.BaseBarony
}

var BDzegh泽格 feud.Barony = &泽格DzeghBarony{}

func init() {
    f := BDzegh泽格.(*泽格DzeghBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dzegh",
		TitleName: "泽格",
		TitleCode: "b_dzegh",
	}
}
