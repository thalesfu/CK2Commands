package mali

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏萨SoussaBarony struct {
	feud.BaseBarony
}

var BSoussa苏萨 feud.Barony = &苏萨SoussaBarony{}

func init() {
    f := BSoussa苏萨.(*苏萨SoussaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "soussa",
		TitleName: "苏萨",
		TitleCode: "b_soussa",
	}
}
