package al_mafraq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马弗拉克MafraqBarony struct {
	feud.BaseBarony
}

var BMafraq马弗拉克 feud.Barony = &马弗拉克MafraqBarony{}

func init() {
    f := BMafraq马弗拉克.(*马弗拉克MafraqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mafraq",
		TitleName: "马弗拉克",
		TitleCode: "b_mafraq",
	}
}
