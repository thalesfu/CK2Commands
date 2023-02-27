package holstein

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔登堡Oldenburg_holsteinBarony struct {
	feud.BaseBarony
}

var BOldenburg_holstein奥尔登堡 feud.Barony = &奥尔登堡Oldenburg_holsteinBarony{}

func init() {
    f := BOldenburg_holstein奥尔登堡.(*奥尔登堡Oldenburg_holsteinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oldenburg_holstein",
		TitleName: "奥尔登堡",
		TitleCode: "b_oldenburg_holstein",
	}
}
