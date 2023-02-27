package nordgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞尔茨SelzBarony struct {
	feud.BaseBarony
}

var BSelz塞尔茨 feud.Barony = &塞尔茨SelzBarony{}

func init() {
    f := BSelz塞尔茨.(*塞尔茨SelzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "selz",
		TitleName: "塞尔茨",
		TitleCode: "b_selz",
	}
}
