package ghazna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加德兹GardezBarony struct {
	feud.BaseBarony
}

var BGardez加德兹 feud.Barony = &加德兹GardezBarony{}

func init() {
    f := BGardez加德兹.(*加德兹GardezBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gardez",
		TitleName: "加德兹",
		TitleCode: "b_gardez",
	}
}
