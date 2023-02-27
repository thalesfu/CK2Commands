package agrigento

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰拉GelaBarony struct {
	feud.BaseBarony
}

var BGela杰拉 feud.Barony = &杰拉GelaBarony{}

func init() {
    f := BGela杰拉.(*杰拉GelaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gela",
		TitleName: "杰拉",
		TitleCode: "b_gela",
	}
}
