package fitri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马努瓦希ManouachiBarony struct {
	feud.BaseBarony
}

var BManouachi马努瓦希 feud.Barony = &马努瓦希ManouachiBarony{}

func init() {
	f := BManouachi马努瓦希.(*马努瓦希ManouachiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "manouachi",
		TitleName: "马努瓦希",
		TitleCode: "b_manouachi",
	}
}
