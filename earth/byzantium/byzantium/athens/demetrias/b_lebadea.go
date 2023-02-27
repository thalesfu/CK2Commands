package demetrias

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱瓦贾LebadeaBarony struct {
	feud.BaseBarony
}

var BLebadea莱瓦贾 feud.Barony = &莱瓦贾LebadeaBarony{}

func init() {
    f := BLebadea莱瓦贾.(*莱瓦贾LebadeaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lebadea",
		TitleName: "莱瓦贾",
		TitleCode: "b_lebadea",
	}
}
