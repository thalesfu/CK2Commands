package maragha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米亚内MiyanehBarony struct {
	feud.BaseBarony
}

var BMiyaneh米亚内 feud.Barony = &米亚内MiyanehBarony{}

func init() {
	f := BMiyaneh米亚内.(*米亚内MiyanehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "miyaneh",
		TitleName: "米亚内",
		TitleCode: "b_miyaneh",
	}
}
