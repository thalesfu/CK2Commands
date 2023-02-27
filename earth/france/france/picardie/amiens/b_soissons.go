package amiens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏瓦松SoissonsBarony struct {
	feud.BaseBarony
}

var BSoissons苏瓦松 feud.Barony = &苏瓦松SoissonsBarony{}

func init() {
    f := BSoissons苏瓦松.(*苏瓦松SoissonsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "soissons",
		TitleName: "苏瓦松",
		TitleCode: "b_soissons",
	}
}
