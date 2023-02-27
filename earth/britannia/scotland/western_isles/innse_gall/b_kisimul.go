package innse_gall

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基西穆尔KisimulBarony struct {
	feud.BaseBarony
}

var BKisimul基西穆尔 feud.Barony = &基西穆尔KisimulBarony{}

func init() {
    f := BKisimul基西穆尔.(*基西穆尔KisimulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kisimul",
		TitleName: "基西穆尔",
		TitleCode: "b_kisimul",
	}
}
