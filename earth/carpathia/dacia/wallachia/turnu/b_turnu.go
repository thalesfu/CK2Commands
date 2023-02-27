package turnu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图尔努TurnuBarony struct {
	feud.BaseBarony
}

var BTurnu图尔努 feud.Barony = &图尔努TurnuBarony{}

func init() {
    f := BTurnu图尔努.(*图尔努TurnuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "turnu",
		TitleName: "图尔努",
		TitleCode: "b_turnu",
	}
}
