package furstenberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索伦ZollernBarony struct {
	feud.BaseBarony
}

var BZollern索伦 feud.Barony = &索伦ZollernBarony{}

func init() {
    f := BZollern索伦.(*索伦ZollernBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zollern",
		TitleName: "索伦",
		TitleCode: "b_zollern",
	}
}
