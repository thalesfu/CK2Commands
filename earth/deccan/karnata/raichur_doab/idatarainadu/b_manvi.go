package idatarainadu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼维ManviBarony struct {
	feud.BaseBarony
}

var BManvi曼维 feud.Barony = &曼维ManviBarony{}

func init() {
    f := BManvi曼维.(*曼维ManviBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "manvi",
		TitleName: "曼维",
		TitleCode: "b_manvi",
	}
}
