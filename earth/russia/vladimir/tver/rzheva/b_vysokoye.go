package rzheva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维索科耶VysokoyeBarony struct {
	feud.BaseBarony
}

var BVysokoye维索科耶 feud.Barony = &维索科耶VysokoyeBarony{}

func init() {
    f := BVysokoye维索科耶.(*维索科耶VysokoyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vysokoye",
		TitleName: "维索科耶",
		TitleCode: "b_vysokoye",
	}
}
