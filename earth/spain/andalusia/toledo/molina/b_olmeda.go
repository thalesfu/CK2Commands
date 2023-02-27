package molina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔梅达OlmedaBarony struct {
	feud.BaseBarony
}

var BOlmeda奥尔梅达 feud.Barony = &奥尔梅达OlmedaBarony{}

func init() {
    f := BOlmeda奥尔梅达.(*奥尔梅达OlmedaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "olmeda",
		TitleName: "奥尔梅达",
		TitleCode: "b_olmeda",
	}
}
