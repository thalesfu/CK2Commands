package amastris

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库托洛斯CyttorosBarony struct {
	feud.BaseBarony
}

var BCyttoros库托洛斯 feud.Barony = &库托洛斯CyttorosBarony{}

func init() {
	f := BCyttoros库托洛斯.(*库托洛斯CyttorosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cyttoros",
		TitleName: "库托洛斯",
		TitleCode: "b_cyttoros",
	}
}
