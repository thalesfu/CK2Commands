package al_aqabah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费法FeifaBarony struct {
	feud.BaseBarony
}

var BFeifa费法 feud.Barony = &费法FeifaBarony{}

func init() {
    f := BFeifa费法.(*费法FeifaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "feifa",
		TitleName: "费法",
		TitleCode: "b_feifa",
	}
}
