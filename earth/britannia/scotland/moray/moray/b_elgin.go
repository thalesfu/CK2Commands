package moray

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃尔金ElginBarony struct {
	feud.BaseBarony
}

var BElgin埃尔金 feud.Barony = &埃尔金ElginBarony{}

func init() {
	f := BElgin埃尔金.(*埃尔金ElginBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elgin",
		TitleName: "埃尔金",
		TitleCode: "b_elgin",
	}
}
