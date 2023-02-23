package thessalia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达玛斯DamasisBarony struct {
	feud.BaseBarony
}

var BDamasis达玛斯 feud.Barony = &达玛斯DamasisBarony{}

func init() {
	f := BDamasis达玛斯.(*达玛斯DamasisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "damasis",
		TitleName: "达玛斯",
		TitleCode: "b_damasis",
	}
}
