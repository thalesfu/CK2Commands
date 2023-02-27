package gwynedd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪加努伊DegannwyBarony struct {
	feud.BaseBarony
}

var BDegannwy迪加努伊 feud.Barony = &迪加努伊DegannwyBarony{}

func init() {
    f := BDegannwy迪加努伊.(*迪加努伊DegannwyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "degannwy",
		TitleName: "迪加努伊",
		TitleCode: "b_degannwy",
	}
}
