package valencia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙蒂瓦JativaBarony struct {
	feud.BaseBarony
}

var BJativa沙蒂瓦 feud.Barony = &沙蒂瓦JativaBarony{}

func init() {
    f := BJativa沙蒂瓦.(*沙蒂瓦JativaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jativa",
		TitleName: "沙蒂瓦",
		TitleCode: "b_jativa",
	}
}
