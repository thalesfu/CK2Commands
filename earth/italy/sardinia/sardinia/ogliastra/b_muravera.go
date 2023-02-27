package ogliastra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆拉韦拉MuraveraBarony struct {
	feud.BaseBarony
}

var BMuravera穆拉韦拉 feud.Barony = &穆拉韦拉MuraveraBarony{}

func init() {
    f := BMuravera穆拉韦拉.(*穆拉韦拉MuraveraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "muravera",
		TitleName: "穆拉韦拉",
		TitleCode: "b_muravera",
	}
}
