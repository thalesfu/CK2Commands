package hastinapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 象城HastinapuraBarony struct {
	feud.BaseBarony
}

var BHastinapura象城 feud.Barony = &象城HastinapuraBarony{}

func init() {
    f := BHastinapura象城.(*象城HastinapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hastinapura",
		TitleName: "象城",
		TitleCode: "b_hastinapura",
	}
}
