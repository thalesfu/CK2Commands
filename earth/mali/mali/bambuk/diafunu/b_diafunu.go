package diafunu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪亚富努DiafunuBarony struct {
	feud.BaseBarony
}

var BDiafunu迪亚富努 feud.Barony = &迪亚富努DiafunuBarony{}

func init() {
	f := BDiafunu迪亚富努.(*迪亚富努DiafunuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "diafunu",
		TitleName: "迪亚富努",
		TitleCode: "b_diafunu",
	}
}
