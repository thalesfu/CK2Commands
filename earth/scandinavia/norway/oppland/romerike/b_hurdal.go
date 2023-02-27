package romerike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 许达尔HurdalBarony struct {
	feud.BaseBarony
}

var BHurdal许达尔 feud.Barony = &许达尔HurdalBarony{}

func init() {
    f := BHurdal许达尔.(*许达尔HurdalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hurdal",
		TitleName: "许达尔",
		TitleCode: "b_hurdal",
	}
}
