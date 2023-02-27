package bidar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿伐瓦迪AvarwadiBarony struct {
	feud.BaseBarony
}

var BAvarwadi阿伐瓦迪 feud.Barony = &阿伐瓦迪AvarwadiBarony{}

func init() {
    f := BAvarwadi阿伐瓦迪.(*阿伐瓦迪AvarwadiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "avarwadi",
		TitleName: "阿伐瓦迪",
		TitleCode: "b_avarwadi",
	}
}
