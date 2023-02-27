package nellore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎杜古拉KandukuraBarony struct {
	feud.BaseBarony
}

var BKandukura坎杜古拉 feud.Barony = &坎杜古拉KandukuraBarony{}

func init() {
    f := BKandukura坎杜古拉.(*坎杜古拉KandukuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kandukura",
		TitleName: "坎杜古拉",
		TitleCode: "b_kandukura",
	}
}
