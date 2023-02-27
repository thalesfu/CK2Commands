package air

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿加德斯AgadesBarony struct {
	feud.BaseBarony
}

var BAgades阿加德斯 feud.Barony = &阿加德斯AgadesBarony{}

func init() {
    f := BAgades阿加德斯.(*阿加德斯AgadesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "agades",
		TitleName: "阿加德斯",
		TitleCode: "b_agades",
	}
}
