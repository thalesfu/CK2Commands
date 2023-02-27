package lothian

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯特灵StirlingBarony struct {
	feud.BaseBarony
}

var BStirling斯特灵 feud.Barony = &斯特灵StirlingBarony{}

func init() {
    f := BStirling斯特灵.(*斯特灵StirlingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stirling",
		TitleName: "斯特灵",
		TitleCode: "b_stirling",
	}
}
