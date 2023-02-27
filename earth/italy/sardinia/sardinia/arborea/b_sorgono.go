package arborea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索尔戈诺SorgonoBarony struct {
	feud.BaseBarony
}

var BSorgono索尔戈诺 feud.Barony = &索尔戈诺SorgonoBarony{}

func init() {
    f := BSorgono索尔戈诺.(*索尔戈诺SorgonoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sorgono",
		TitleName: "索尔戈诺",
		TitleCode: "b_sorgono",
	}
}
