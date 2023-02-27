package roma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏特里SutriBarony struct {
	feud.BaseBarony
}

var BSutri苏特里 feud.Barony = &苏特里SutriBarony{}

func init() {
    f := BSutri苏特里.(*苏特里SutriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sutri",
		TitleName: "苏特里",
		TitleCode: "b_sutri",
	}
}
