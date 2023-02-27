package hendjan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏萨SusaBarony struct {
	feud.BaseBarony
}

var BSusa苏萨 feud.Barony = &苏萨SusaBarony{}

func init() {
    f := BSusa苏萨.(*苏萨SusaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "susa",
		TitleName: "苏萨",
		TitleCode: "b_susa",
	}
}
