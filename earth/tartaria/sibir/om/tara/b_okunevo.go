package tara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥库涅沃OkunevoBarony struct {
	feud.BaseBarony
}

var BOkunevo奥库涅沃 feud.Barony = &奥库涅沃OkunevoBarony{}

func init() {
    f := BOkunevo奥库涅沃.(*奥库涅沃OkunevoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "okunevo",
		TitleName: "奥库涅沃",
		TitleCode: "b_okunevo",
	}
}
