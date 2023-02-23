package zadar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 本科瓦茨BenkovacBarony struct {
	feud.BaseBarony
}

var BBenkovac本科瓦茨 feud.Barony = &本科瓦茨BenkovacBarony{}

func init() {
	f := BBenkovac本科瓦茨.(*本科瓦茨BenkovacBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "benkovac",
		TitleName: "本科瓦茨",
		TitleCode: "b_benkovac",
	}
}
