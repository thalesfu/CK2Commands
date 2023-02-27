package gotland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔瓦AlvaBarony struct {
	feud.BaseBarony
}

var BAlva阿尔瓦 feud.Barony = &阿尔瓦AlvaBarony{}

func init() {
    f := BAlva阿尔瓦.(*阿尔瓦AlvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alva",
		TitleName: "阿尔瓦",
		TitleCode: "b_alva",
	}
}
