package markam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莽岭ManglingBarony struct {
	feud.BaseBarony
}

var BMangling莽岭 feud.Barony = &莽岭ManglingBarony{}

func init() {
    f := BMangling莽岭.(*莽岭ManglingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mangling",
		TitleName: "莽岭",
		TitleCode: "b_mangling",
	}
}
