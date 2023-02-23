package lucca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维亚雷焦ViareggioBarony struct {
	feud.BaseBarony
}

var BViareggio维亚雷焦 feud.Barony = &维亚雷焦ViareggioBarony{}

func init() {
	f := BViareggio维亚雷焦.(*维亚雷焦ViareggioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "viareggio",
		TitleName: "维亚雷焦",
		TitleCode: "b_viareggio",
	}
}
