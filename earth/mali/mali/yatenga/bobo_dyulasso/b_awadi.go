package bobo_dyulasso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿瓦迪AwadiBarony struct {
	feud.BaseBarony
}

var BAwadi阿瓦迪 feud.Barony = &阿瓦迪AwadiBarony{}

func init() {
    f := BAwadi阿瓦迪.(*阿瓦迪AwadiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "awadi",
		TitleName: "阿瓦迪",
		TitleCode: "b_awadi",
	}
}
