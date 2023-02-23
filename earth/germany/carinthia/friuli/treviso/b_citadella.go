package treviso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奇塔代拉CitadellaBarony struct {
	feud.BaseBarony
}

var BCitadella奇塔代拉 feud.Barony = &奇塔代拉CitadellaBarony{}

func init() {
	f := BCitadella奇塔代拉.(*奇塔代拉CitadellaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "citadella",
		TitleName: "奇塔代拉",
		TitleCode: "b_citadella",
	}
}
