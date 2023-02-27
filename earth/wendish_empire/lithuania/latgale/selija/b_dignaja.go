package selija

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪格纳亚DignajaBarony struct {
	feud.BaseBarony
}

var BDignaja迪格纳亚 feud.Barony = &迪格纳亚DignajaBarony{}

func init() {
    f := BDignaja迪格纳亚.(*迪格纳亚DignajaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dignaja",
		TitleName: "迪格纳亚",
		TitleCode: "b_dignaja",
	}
}
