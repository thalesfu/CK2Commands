package chios

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 富尔尼FourniBarony struct {
	feud.BaseBarony
}

var BFourni富尔尼 feud.Barony = &富尔尼FourniBarony{}

func init() {
    f := BFourni富尔尼.(*富尔尼FourniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fourni",
		TitleName: "富尔尼",
		TitleCode: "b_fourni",
	}
}
