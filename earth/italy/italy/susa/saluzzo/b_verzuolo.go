package saluzzo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦尔佐洛VerzuoloBarony struct {
	feud.BaseBarony
}

var BVerzuolo韦尔佐洛 feud.Barony = &韦尔佐洛VerzuoloBarony{}

func init() {
    f := BVerzuolo韦尔佐洛.(*韦尔佐洛VerzuoloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "verzuolo",
		TitleName: "韦尔佐洛",
		TitleCode: "b_verzuolo",
	}
}
