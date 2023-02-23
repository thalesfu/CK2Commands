package bath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦尔斯WellsBarony struct {
	feud.BaseBarony
}

var BWells韦尔斯 feud.Barony = &韦尔斯WellsBarony{}

func init() {
	f := BWells韦尔斯.(*韦尔斯WellsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wells",
		TitleName: "韦尔斯",
		TitleCode: "b_wells",
	}
}
