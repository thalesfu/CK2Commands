package qarazhyrya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔套OrtauBarony struct {
	feud.BaseBarony
}

var BOrtau奥尔套 feud.Barony = &奥尔套OrtauBarony{}

func init() {
	f := BOrtau奥尔套.(*奥尔套OrtauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ortau",
		TitleName: "奥尔套",
		TitleCode: "b_ortau",
	}
}
