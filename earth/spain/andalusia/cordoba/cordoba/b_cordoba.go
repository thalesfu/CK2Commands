package cordoba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔多瓦CordobaBarony struct {
	feud.BaseBarony
}

var BCordoba科尔多瓦 feud.Barony = &科尔多瓦CordobaBarony{}

func init() {
    f := BCordoba科尔多瓦.(*科尔多瓦CordobaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cordoba",
		TitleName: "科尔多瓦",
		TitleCode: "b_cordoba",
	}
}
