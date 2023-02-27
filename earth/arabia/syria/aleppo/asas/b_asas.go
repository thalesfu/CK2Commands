package asas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿萨斯AsasBarony struct {
	feud.BaseBarony
}

var BAsas阿萨斯 feud.Barony = &阿萨斯AsasBarony{}

func init() {
    f := BAsas阿萨斯.(*阿萨斯AsasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asas",
		TitleName: "阿萨斯",
		TitleCode: "b_asas",
	}
}
