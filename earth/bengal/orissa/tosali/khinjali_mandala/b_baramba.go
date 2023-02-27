package khinjali_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴兰巴BarambaBarony struct {
	feud.BaseBarony
}

var BBaramba巴兰巴 feud.Barony = &巴兰巴BarambaBarony{}

func init() {
    f := BBaramba巴兰巴.(*巴兰巴BarambaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baramba",
		TitleName: "巴兰巴",
		TitleCode: "b_baramba",
	}
}
