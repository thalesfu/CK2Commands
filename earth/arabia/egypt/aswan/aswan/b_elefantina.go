package aswan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 象岛ElefantinaBarony struct {
	feud.BaseBarony
}

var BElefantina象岛 feud.Barony = &象岛ElefantinaBarony{}

func init() {
    f := BElefantina象岛.(*象岛ElefantinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elefantina",
		TitleName: "象岛",
		TitleCode: "b_elefantina",
	}
}
