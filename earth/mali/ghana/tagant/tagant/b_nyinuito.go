package tagant

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼努伊托NyinuitoBarony struct {
	feud.BaseBarony
}

var BNyinuito尼努伊托 feud.Barony = &尼努伊托NyinuitoBarony{}

func init() {
	f := BNyinuito尼努伊托.(*尼努伊托NyinuitoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nyinuito",
		TitleName: "尼努伊托",
		TitleCode: "b_nyinuito",
	}
}
