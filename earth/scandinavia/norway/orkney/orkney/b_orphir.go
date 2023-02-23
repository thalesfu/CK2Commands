package orkney

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥费尔OrphirBarony struct {
	feud.BaseBarony
}

var BOrphir奥费尔 feud.Barony = &奥费尔OrphirBarony{}

func init() {
	f := BOrphir奥费尔.(*奥费尔OrphirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "orphir",
		TitleName: "奥费尔",
		TitleCode: "b_orphir",
	}
}
