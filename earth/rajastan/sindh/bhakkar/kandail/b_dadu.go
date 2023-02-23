package kandail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达都DaduBarony struct {
	feud.BaseBarony
}

var BDadu达都 feud.Barony = &达都DaduBarony{}

func init() {
	f := BDadu达都.(*达都DaduBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dadu",
		TitleName: "达都",
		TitleCode: "b_dadu",
	}
}
