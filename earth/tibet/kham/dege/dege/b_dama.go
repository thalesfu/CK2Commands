package dege

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达马DamaBarony struct {
	feud.BaseBarony
}

var BDama达马 feud.Barony = &达马DamaBarony{}

func init() {
	f := BDama达马.(*达马DamaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dama",
		TitleName: "达马",
		TitleCode: "b_dama",
	}
}
