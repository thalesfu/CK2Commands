package khuiten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索尔韦诺克SorvenokBarony struct {
	feud.BaseBarony
}

var BSorvenok索尔韦诺克 feud.Barony = &索尔韦诺克SorvenokBarony{}

func init() {
	f := BSorvenok索尔韦诺克.(*索尔韦诺克SorvenokBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sorvenok",
		TitleName: "索尔韦诺克",
		TitleCode: "b_sorvenok",
	}
}
