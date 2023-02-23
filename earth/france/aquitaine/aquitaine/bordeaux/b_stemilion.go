package bordeaux

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣埃米利翁StemilionBarony struct {
	feud.BaseBarony
}

var BStemilion圣埃米利翁 feud.Barony = &圣埃米利翁StemilionBarony{}

func init() {
	f := BStemilion圣埃米利翁.(*圣埃米利翁StemilionBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stemilion",
		TitleName: "圣埃米利翁",
		TitleCode: "b_stemilion",
	}
}
