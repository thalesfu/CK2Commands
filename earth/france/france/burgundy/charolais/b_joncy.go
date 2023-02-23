package charolais

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 容西JoncyBarony struct {
	feud.BaseBarony
}

var BJoncy容西 feud.Barony = &容西JoncyBarony{}

func init() {
	f := BJoncy容西.(*容西JoncyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "joncy",
		TitleName: "容西",
		TitleCode: "b_joncy",
	}
}
