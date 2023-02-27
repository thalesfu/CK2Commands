package armagnac

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱克图尔LectoureBarony struct {
	feud.BaseBarony
}

var BLectoure莱克图尔 feud.Barony = &莱克图尔LectoureBarony{}

func init() {
    f := BLectoure莱克图尔.(*莱克图尔LectoureBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lectoure",
		TitleName: "莱克图尔",
		TitleCode: "b_lectoure",
	}
}
