package lleida

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱斯博尔热斯布兰克斯BorgesblanquesBarony struct {
	feud.BaseBarony
}

var BBorgesblanques莱斯博尔热斯布兰克斯 feud.Barony = &莱斯博尔热斯布兰克斯BorgesblanquesBarony{}

func init() {
    f := BBorgesblanques莱斯博尔热斯布兰克斯.(*莱斯博尔热斯布兰克斯BorgesblanquesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "borgesblanques",
		TitleName: "莱斯博尔热斯布兰克斯",
		TitleCode: "b_borgesblanques",
	}
}
