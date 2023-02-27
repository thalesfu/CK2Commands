package upper_silesia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥莱希尼察OlesnicaBarony struct {
	feud.BaseBarony
}

var BOlesnica奥莱希尼察 feud.Barony = &奥莱希尼察OlesnicaBarony{}

func init() {
    f := BOlesnica奥莱希尼察.(*奥莱希尼察OlesnicaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "olesnica",
		TitleName: "奥莱希尼察",
		TitleCode: "b_olesnica",
	}
}
