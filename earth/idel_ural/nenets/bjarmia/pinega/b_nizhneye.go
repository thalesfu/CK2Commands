package pinega

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼日涅耶NizhneyeBarony struct {
	feud.BaseBarony
}

var BNizhneye尼日涅耶 feud.Barony = &尼日涅耶NizhneyeBarony{}

func init() {
    f := BNizhneye尼日涅耶.(*尼日涅耶NizhneyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nizhneye",
		TitleName: "尼日涅耶",
		TitleCode: "b_nizhneye",
	}
}
