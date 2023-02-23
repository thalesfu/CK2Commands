package astorga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法韦罗FaberoBarony struct {
	feud.BaseBarony
}

var BFabero法韦罗 feud.Barony = &法韦罗FaberoBarony{}

func init() {
	f := BFabero法韦罗.(*法韦罗FaberoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fabero",
		TitleName: "法韦罗",
		TitleCode: "b_fabero",
	}
}
