package kyzikos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮尔科特PercoteBarony struct {
	feud.BaseBarony
}

var BPercote皮尔科特 feud.Barony = &皮尔科特PercoteBarony{}

func init() {
    f := BPercote皮尔科特.(*皮尔科特PercoteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "percote",
		TitleName: "皮尔科特",
		TitleCode: "b_percote",
	}
}
