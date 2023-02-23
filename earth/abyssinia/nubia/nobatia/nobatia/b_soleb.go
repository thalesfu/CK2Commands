package nobatia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 叟勒布SolebBarony struct {
	feud.BaseBarony
}

var BSoleb叟勒布 feud.Barony = &叟勒布SolebBarony{}

func init() {
	f := BSoleb叟勒布.(*叟勒布SolebBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "soleb",
		TitleName: "叟勒布",
		TitleCode: "b_soleb",
	}
}
