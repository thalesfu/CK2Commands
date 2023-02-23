package sinope

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔劳拉TalauraBarony struct {
	feud.BaseBarony
}

var BTalaura塔劳拉 feud.Barony = &塔劳拉TalauraBarony{}

func init() {
	f := BTalaura塔劳拉.(*塔劳拉TalauraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "talaura",
		TitleName: "塔劳拉",
		TitleCode: "b_talaura",
	}
}
