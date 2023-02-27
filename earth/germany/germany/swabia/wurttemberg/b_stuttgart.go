package wurttemberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯图加特StuttgartBarony struct {
	feud.BaseBarony
}

var BStuttgart斯图加特 feud.Barony = &斯图加特StuttgartBarony{}

func init() {
    f := BStuttgart斯图加特.(*斯图加特StuttgartBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stuttgart",
		TitleName: "斯图加特",
		TitleCode: "b_stuttgart",
	}
}
