package zadar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥布罗瓦茨ObrovacBarony struct {
	feud.BaseBarony
}

var BObrovac奥布罗瓦茨 feud.Barony = &奥布罗瓦茨ObrovacBarony{}

func init() {
	f := BObrovac奥布罗瓦茨.(*奥布罗瓦茨ObrovacBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "obrovac",
		TitleName: "奥布罗瓦茨",
		TitleCode: "b_obrovac",
	}
}
