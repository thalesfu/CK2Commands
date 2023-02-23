package manupura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞易斯SaisBarony struct {
	feud.BaseBarony
}

var BSais塞易斯 feud.Barony = &塞易斯SaisBarony{}

func init() {
	f := BSais塞易斯.(*塞易斯SaisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sais",
		TitleName: "塞易斯",
		TitleCode: "b_sais",
	}
}
