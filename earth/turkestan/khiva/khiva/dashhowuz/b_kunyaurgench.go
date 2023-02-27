package dashhowuz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古玉龙杰赤KunyaurgenchBarony struct {
	feud.BaseBarony
}

var BKunyaurgench古玉龙杰赤 feud.Barony = &古玉龙杰赤KunyaurgenchBarony{}

func init() {
    f := BKunyaurgench古玉龙杰赤.(*古玉龙杰赤KunyaurgenchBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kunyaurgench",
		TitleName: "古玉龙杰赤",
		TitleCode: "b_kunyaurgench",
	}
}
