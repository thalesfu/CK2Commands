package tadjoura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪基勒DikhilBarony struct {
	feud.BaseBarony
}

var BDikhil迪基勒 feud.Barony = &迪基勒DikhilBarony{}

func init() {
    f := BDikhil迪基勒.(*迪基勒DikhilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dikhil",
		TitleName: "迪基勒",
		TitleCode: "b_dikhil",
	}
}
