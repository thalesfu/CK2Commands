package la_marche

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖雷GueretBarony struct {
	feud.BaseBarony
}

var BGueret盖雷 feud.Barony = &盖雷GueretBarony{}

func init() {
    f := BGueret盖雷.(*盖雷GueretBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gueret",
		TitleName: "盖雷",
		TitleCode: "b_gueret",
	}
}
