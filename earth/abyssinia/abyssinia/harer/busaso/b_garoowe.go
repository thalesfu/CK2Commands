package busaso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加罗韦GarooweBarony struct {
	feud.BaseBarony
}

var BGaroowe加罗韦 feud.Barony = &加罗韦GarooweBarony{}

func init() {
	f := BGaroowe加罗韦.(*加罗韦GarooweBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "garoowe",
		TitleName: "加罗韦",
		TitleCode: "b_garoowe",
	}
}
