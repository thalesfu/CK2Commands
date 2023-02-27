package rovaniemi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 内凯莱NakkalaBarony struct {
	feud.BaseBarony
}

var BNakkala内凯莱 feud.Barony = &内凯莱NakkalaBarony{}

func init() {
    f := BNakkala内凯莱.(*内凯莱NakkalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nakkala",
		TitleName: "内凯莱",
		TitleCode: "b_nakkala",
	}
}
