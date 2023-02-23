package valladolid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿维拉AvilaBarony struct {
	feud.BaseBarony
}

var BAvila阿维拉 feud.Barony = &阿维拉AvilaBarony{}

func init() {
	f := BAvila阿维拉.(*阿维拉AvilaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "avila",
		TitleName: "阿维拉",
		TitleCode: "b_avila",
	}
}
