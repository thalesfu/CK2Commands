package avhaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪兹富勒DezfulBarony struct {
	feud.BaseBarony
}

var BDezful迪兹富勒 feud.Barony = &迪兹富勒DezfulBarony{}

func init() {
    f := BDezful迪兹富勒.(*迪兹富勒DezfulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dezful",
		TitleName: "迪兹富勒",
		TitleCode: "b_dezful",
	}
}
