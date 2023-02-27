package famagusta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣伊拉里SthilarionBarony struct {
	feud.BaseBarony
}

var BSthilarion圣伊拉里 feud.Barony = &圣伊拉里SthilarionBarony{}

func init() {
    f := BSthilarion圣伊拉里.(*圣伊拉里SthilarionBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sthilarion",
		TitleName: "圣伊拉里",
		TitleCode: "b_sthilarion",
	}
}
