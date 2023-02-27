package mahra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼什通NishtunBarony struct {
	feud.BaseBarony
}

var BNishtun尼什通 feud.Barony = &尼什通NishtunBarony{}

func init() {
    f := BNishtun尼什通.(*尼什通NishtunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nishtun",
		TitleName: "尼什通",
		TitleCode: "b_nishtun",
	}
}
