package palencia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔梅达Olmeda_palenciaBarony struct {
	feud.BaseBarony
}

var BOlmeda_palencia奥尔梅达 feud.Barony = &奥尔梅达Olmeda_palenciaBarony{}

func init() {
    f := BOlmeda_palencia奥尔梅达.(*奥尔梅达Olmeda_palenciaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "olmeda_palencia",
		TitleName: "奥尔梅达",
		TitleCode: "b_olmeda_palencia",
	}
}
