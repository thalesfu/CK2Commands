package narbonne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳博讷NarbonneBarony struct {
	feud.BaseBarony
}

var BNarbonne纳博讷 feud.Barony = &纳博讷NarbonneBarony{}

func init() {
    f := BNarbonne纳博讷.(*纳博讷NarbonneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "narbonne",
		TitleName: "纳博讷",
		TitleCode: "b_narbonne",
	}
}
