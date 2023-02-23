package gao

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 昂沙瓦迪AnchawadiBarony struct {
	feud.BaseBarony
}

var BAnchawadi昂沙瓦迪 feud.Barony = &昂沙瓦迪AnchawadiBarony{}

func init() {
	f := BAnchawadi昂沙瓦迪.(*昂沙瓦迪AnchawadiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anchawadi",
		TitleName: "昂沙瓦迪",
		TitleCode: "b_anchawadi",
	}
}
