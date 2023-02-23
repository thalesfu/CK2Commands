package pettau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 魏滕施泰因WeitensteinBarony struct {
	feud.BaseBarony
}

var BWeitenstein魏滕施泰因 feud.Barony = &魏滕施泰因WeitensteinBarony{}

func init() {
	f := BWeitenstein魏滕施泰因.(*魏滕施泰因WeitensteinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "weitenstein",
		TitleName: "魏滕施泰因",
		TitleCode: "b_weitenstein",
	}
}
