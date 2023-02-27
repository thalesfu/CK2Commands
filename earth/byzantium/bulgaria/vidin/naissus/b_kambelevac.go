package naissus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎贝莱瓦茨KambelevacBarony struct {
	feud.BaseBarony
}

var BKambelevac坎贝莱瓦茨 feud.Barony = &坎贝莱瓦茨KambelevacBarony{}

func init() {
    f := BKambelevac坎贝莱瓦茨.(*坎贝莱瓦茨KambelevacBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kambelevac",
		TitleName: "坎贝莱瓦茨",
		TitleCode: "b_kambelevac",
	}
}
