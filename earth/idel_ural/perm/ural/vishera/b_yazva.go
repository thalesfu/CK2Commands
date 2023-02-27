package vishera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚济瓦YazvaBarony struct {
	feud.BaseBarony
}

var BYazva亚济瓦 feud.Barony = &亚济瓦YazvaBarony{}

func init() {
    f := BYazva亚济瓦.(*亚济瓦YazvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yazva",
		TitleName: "亚济瓦",
		TitleCode: "b_yazva",
	}
}
