package hull

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍恩西HornseaBarony struct {
	feud.BaseBarony
}

var BHornsea霍恩西 feud.Barony = &霍恩西HornseaBarony{}

func init() {
	f := BHornsea霍恩西.(*霍恩西HornseaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hornsea",
		TitleName: "霍恩西",
		TitleCode: "b_hornsea",
	}
}
