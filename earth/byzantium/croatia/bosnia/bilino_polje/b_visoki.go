package bilino_polje

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维索基VisokiBarony struct {
	feud.BaseBarony
}

var BVisoki维索基 feud.Barony = &维索基VisokiBarony{}

func init() {
    f := BVisoki维索基.(*维索基VisokiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "visoki",
		TitleName: "维索基",
		TitleCode: "b_visoki",
	}
}
