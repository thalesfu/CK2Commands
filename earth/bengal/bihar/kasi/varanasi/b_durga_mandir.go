package varanasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 突伽女神庙Durga_mandirBarony struct {
	feud.BaseBarony
}

var BDurga_mandir突伽女神庙 feud.Barony = &突伽女神庙Durga_mandirBarony{}

func init() {
    f := BDurga_mandir突伽女神庙.(*突伽女神庙Durga_mandirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "durga_mandir",
		TitleName: "突伽女神庙",
		TitleCode: "b_durga_mandir",
	}
}
