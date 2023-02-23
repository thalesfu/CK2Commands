package sodermanland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼雪平NykopingBarony struct {
	feud.BaseBarony
}

var BNykoping尼雪平 feud.Barony = &尼雪平NykopingBarony{}

func init() {
	f := BNykoping尼雪平.(*尼雪平NykopingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nykoping",
		TitleName: "尼雪平",
		TitleCode: "b_nykoping",
	}
}
