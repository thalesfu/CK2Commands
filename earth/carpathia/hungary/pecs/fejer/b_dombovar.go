package fejer

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 栋博堡DombovarBarony struct {
	feud.BaseBarony
}

var BDombovar栋博堡 feud.Barony = &栋博堡DombovarBarony{}

func init() {
	f := BDombovar栋博堡.(*栋博堡DombovarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dombovar",
		TitleName: "栋博堡",
		TitleCode: "b_dombovar",
	}
}
