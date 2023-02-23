package bourbon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维希VichyBarony struct {
	feud.BaseBarony
}

var BVichy维希 feud.Barony = &维希VichyBarony{}

func init() {
	f := BVichy维希.(*维希VichyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vichy",
		TitleName: "维希",
		TitleCode: "b_vichy",
	}
}
