package tsakha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别若则PerotseBarony struct {
	feud.BaseBarony
}

var BPerotse别若则 feud.Barony = &别若则PerotseBarony{}

func init() {
	f := BPerotse别若则.(*别若则PerotseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "perotse",
		TitleName: "别若则",
		TitleCode: "b_perotse",
	}
}
