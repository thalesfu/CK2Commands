package tus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索拉克SolakBarony struct {
	feud.BaseBarony
}

var BSolak索拉克 feud.Barony = &索拉克SolakBarony{}

func init() {
	f := BSolak索拉克.(*索拉克SolakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "solak",
		TitleName: "索拉克",
		TitleCode: "b_solak",
	}
}
