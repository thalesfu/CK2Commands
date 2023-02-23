package onogost

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格拉迪纳GradinaBarony struct {
	feud.BaseBarony
}

var BGradina格拉迪纳 feud.Barony = &格拉迪纳GradinaBarony{}

func init() {
	f := BGradina格拉迪纳.(*格拉迪纳GradinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gradina",
		TitleName: "格拉迪纳",
		TitleCode: "b_gradina",
	}
}
