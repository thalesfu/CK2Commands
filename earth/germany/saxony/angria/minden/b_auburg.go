package minden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥堡AuburgBarony struct {
	feud.BaseBarony
}

var BAuburg奥堡 feud.Barony = &奥堡AuburgBarony{}

func init() {
	f := BAuburg奥堡.(*奥堡AuburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "auburg",
		TitleName: "奥堡",
		TitleCode: "b_auburg",
	}
}
