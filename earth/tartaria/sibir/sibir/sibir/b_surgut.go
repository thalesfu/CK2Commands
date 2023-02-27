package sibir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏尔古特SurgutBarony struct {
	feud.BaseBarony
}

var BSurgut苏尔古特 feud.Barony = &苏尔古特SurgutBarony{}

func init() {
    f := BSurgut苏尔古特.(*苏尔古特SurgutBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "surgut",
		TitleName: "苏尔古特",
		TitleCode: "b_surgut",
	}
}
