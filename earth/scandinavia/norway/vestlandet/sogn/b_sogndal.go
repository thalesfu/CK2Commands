package sogn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 松达尔SogndalBarony struct {
	feud.BaseBarony
}

var BSogndal松达尔 feud.Barony = &松达尔SogndalBarony{}

func init() {
	f := BSogndal松达尔.(*松达尔SogndalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sogndal",
		TitleName: "松达尔",
		TitleCode: "b_sogndal",
	}
}
