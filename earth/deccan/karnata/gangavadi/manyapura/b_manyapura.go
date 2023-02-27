package manyapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼耶补罗ManyapuraBarony struct {
	feud.BaseBarony
}

var BManyapura曼耶补罗 feud.Barony = &曼耶补罗ManyapuraBarony{}

func init() {
    f := BManyapura曼耶补罗.(*曼耶补罗ManyapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "manyapura",
		TitleName: "曼耶补罗",
		TitleCode: "b_manyapura",
	}
}
