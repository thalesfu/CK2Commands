package cadiz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣费尔南多SanfernandoBarony struct {
	feud.BaseBarony
}

var BSanfernando圣费尔南多 feud.Barony = &圣费尔南多SanfernandoBarony{}

func init() {
	f := BSanfernando圣费尔南多.(*圣费尔南多SanfernandoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanfernando",
		TitleName: "圣费尔南多",
		TitleCode: "b_sanfernando",
	}
}
