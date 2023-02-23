package viraja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索罗SoroBarony struct {
	feud.BaseBarony
}

var BSoro索罗 feud.Barony = &索罗SoroBarony{}

func init() {
	f := BSoro索罗.(*索罗SoroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "soro",
		TitleName: "索罗",
		TitleCode: "b_soro",
	}
}
