package ceredigion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰彼得LampeterBarony struct {
	feud.BaseBarony
}

var BLampeter兰彼得 feud.Barony = &兰彼得LampeterBarony{}

func init() {
	f := BLampeter兰彼得.(*兰彼得LampeterBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lampeter",
		TitleName: "兰彼得",
		TitleCode: "b_lampeter",
	}
}
