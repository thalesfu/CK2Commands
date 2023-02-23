package lecce

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布林迪西BrindisiBarony struct {
	feud.BaseBarony
}

var BBrindisi布林迪西 feud.Barony = &布林迪西BrindisiBarony{}

func init() {
	f := BBrindisi布林迪西.(*布林迪西BrindisiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "brindisi",
		TitleName: "布林迪西",
		TitleCode: "b_brindisi",
	}
}
