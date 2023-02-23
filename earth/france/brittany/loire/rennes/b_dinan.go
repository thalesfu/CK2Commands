package rennes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪南DinanBarony struct {
	feud.BaseBarony
}

var BDinan迪南 feud.Barony = &迪南DinanBarony{}

func init() {
	f := BDinan迪南.(*迪南DinanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dinan",
		TitleName: "迪南",
		TitleCode: "b_dinan",
	}
}
