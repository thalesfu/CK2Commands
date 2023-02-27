package osnabruck

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅彭MeppenBarony struct {
	feud.BaseBarony
}

var BMeppen梅彭 feud.Barony = &梅彭MeppenBarony{}

func init() {
    f := BMeppen梅彭.(*梅彭MeppenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "meppen",
		TitleName: "梅彭",
		TitleCode: "b_meppen",
	}
}
