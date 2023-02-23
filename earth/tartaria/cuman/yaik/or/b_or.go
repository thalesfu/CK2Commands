package or

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔OrBarony struct {
	feud.BaseBarony
}

var BOr奥尔 feud.Barony = &奥尔OrBarony{}

func init() {
	f := BOr奥尔.(*奥尔OrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "or",
		TitleName: "奥尔",
		TitleCode: "b_or",
	}
}
