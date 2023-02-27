package sirjan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿巴尔古AbarkuhBarony struct {
	feud.BaseBarony
}

var BAbarkuh阿巴尔古 feud.Barony = &阿巴尔古AbarkuhBarony{}

func init() {
    f := BAbarkuh阿巴尔古.(*阿巴尔古AbarkuhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abarkuh",
		TitleName: "阿巴尔古",
		TitleCode: "b_abarkuh",
	}
}
