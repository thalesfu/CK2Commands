package vladimir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈罗霍韦茨GorokhovetsBarony struct {
	feud.BaseBarony
}

var BGorokhovets戈罗霍韦茨 feud.Barony = &戈罗霍韦茨GorokhovetsBarony{}

func init() {
    f := BGorokhovets戈罗霍韦茨.(*戈罗霍韦茨GorokhovetsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gorokhovets",
		TitleName: "戈罗霍韦茨",
		TitleCode: "b_gorokhovets",
	}
}
