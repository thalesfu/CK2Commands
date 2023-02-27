package penugonda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱帕科什LepakshiBarony struct {
	feud.BaseBarony
}

var BLepakshi莱帕科什 feud.Barony = &莱帕科什LepakshiBarony{}

func init() {
    f := BLepakshi莱帕科什.(*莱帕科什LepakshiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lepakshi",
		TitleName: "莱帕科什",
		TitleCode: "b_lepakshi",
	}
}
