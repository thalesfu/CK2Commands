package tangiers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾西拉AsilahBarony struct {
	feud.BaseBarony
}

var BAsilah艾西拉 feud.Barony = &艾西拉AsilahBarony{}

func init() {
    f := BAsilah艾西拉.(*艾西拉AsilahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asilah",
		TitleName: "艾西拉",
		TitleCode: "b_asilah",
	}
}
