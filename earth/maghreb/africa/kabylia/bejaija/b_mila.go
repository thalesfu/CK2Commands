package bejaija

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 密拉MilaBarony struct {
	feud.BaseBarony
}

var BMila密拉 feud.Barony = &密拉MilaBarony{}

func init() {
	f := BMila密拉.(*密拉MilaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mila",
		TitleName: "密拉",
		TitleCode: "b_mila",
	}
}
