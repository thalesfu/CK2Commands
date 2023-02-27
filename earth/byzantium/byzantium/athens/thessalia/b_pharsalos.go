package thessalia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法尔萨洛斯PharsalosBarony struct {
	feud.BaseBarony
}

var BPharsalos法尔萨洛斯 feud.Barony = &法尔萨洛斯PharsalosBarony{}

func init() {
    f := BPharsalos法尔萨洛斯.(*法尔萨洛斯PharsalosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pharsalos",
		TitleName: "法尔萨洛斯",
		TitleCode: "b_pharsalos",
	}
}
