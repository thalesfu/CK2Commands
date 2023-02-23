package brugge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼乌波特NieuwpoortBarony struct {
	feud.BaseBarony
}

var BNieuwpoort尼乌波特 feud.Barony = &尼乌波特NieuwpoortBarony{}

func init() {
	f := BNieuwpoort尼乌波特.(*尼乌波特NieuwpoortBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nieuwpoort",
		TitleName: "尼乌波特",
		TitleCode: "b_nieuwpoort",
	}
}
