package murom

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅连基MelenkiBarony struct {
	feud.BaseBarony
}

var BMelenki梅连基 feud.Barony = &梅连基MelenkiBarony{}

func init() {
    f := BMelenki梅连基.(*梅连基MelenkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "melenki",
		TitleName: "梅连基",
		TitleCode: "b_melenki",
	}
}
