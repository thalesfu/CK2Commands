package krizevci

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 久尔杰瓦茨DurdevacBarony struct {
	feud.BaseBarony
}

var BDurdevac久尔杰瓦茨 feud.Barony = &久尔杰瓦茨DurdevacBarony{}

func init() {
    f := BDurdevac久尔杰瓦茨.(*久尔杰瓦茨DurdevacBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "durdevac",
		TitleName: "久尔杰瓦茨",
		TitleCode: "b_durdevac",
	}
}
