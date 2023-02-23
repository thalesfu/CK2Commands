package lombardia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛迪LodiBarony struct {
	feud.BaseBarony
}

var BLodi洛迪 feud.Barony = &洛迪LodiBarony{}

func init() {
	f := BLodi洛迪.(*洛迪LodiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lodi",
		TitleName: "洛迪",
		TitleCode: "b_lodi",
	}
}
