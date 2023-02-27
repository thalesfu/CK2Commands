package praha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯塔雷梅斯托Stare_mestoBarony struct {
	feud.BaseBarony
}

var BStare_mesto斯塔雷梅斯托 feud.Barony = &斯塔雷梅斯托Stare_mestoBarony{}

func init() {
    f := BStare_mesto斯塔雷梅斯托.(*斯塔雷梅斯托Stare_mestoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stare_mesto",
		TitleName: "斯塔雷梅斯托",
		TitleCode: "b_stare_mesto",
	}
}
