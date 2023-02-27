package vanema

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 文茨皮尔斯VentspilsBarony struct {
	feud.BaseBarony
}

var BVentspils文茨皮尔斯 feud.Barony = &文茨皮尔斯VentspilsBarony{}

func init() {
    f := BVentspils文茨皮尔斯.(*文茨皮尔斯VentspilsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ventspils",
		TitleName: "文茨皮尔斯",
		TitleCode: "b_ventspils",
	}
}
