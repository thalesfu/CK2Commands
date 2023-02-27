package dalarna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊德勒IdreBarony struct {
	feud.BaseBarony
}

var BIdre伊德勒 feud.Barony = &伊德勒IdreBarony{}

func init() {
    f := BIdre伊德勒.(*伊德勒IdreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "idre",
		TitleName: "伊德勒",
		TitleCode: "b_idre",
	}
}
