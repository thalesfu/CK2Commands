package mesopotamia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利杰LiceBarony struct {
	feud.BaseBarony
}

var BLice利杰 feud.Barony = &利杰LiceBarony{}

func init() {
	f := BLice利杰.(*利杰LiceBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lice",
		TitleName: "利杰",
		TitleCode: "b_lice",
	}
}
