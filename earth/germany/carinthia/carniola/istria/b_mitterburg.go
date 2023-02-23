package istria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米特堡MitterburgBarony struct {
	feud.BaseBarony
}

var BMitterburg米特堡 feud.Barony = &米特堡MitterburgBarony{}

func init() {
	f := BMitterburg米特堡.(*米特堡MitterburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mitterburg",
		TitleName: "米特堡",
		TitleCode: "b_mitterburg",
	}
}
