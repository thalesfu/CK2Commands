package yeruslan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔达OrdaBarony struct {
	feud.BaseBarony
}

var BOrda奥尔达 feud.Barony = &奥尔达OrdaBarony{}

func init() {
    f := BOrda奥尔达.(*奥尔达OrdaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "orda",
		TitleName: "奥尔达",
		TitleCode: "b_orda",
	}
}
