package tregor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普卢里沃PlourivoBarony struct {
	feud.BaseBarony
}

var BPlourivo普卢里沃 feud.Barony = &普卢里沃PlourivoBarony{}

func init() {
	f := BPlourivo普卢里沃.(*普卢里沃PlourivoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "plourivo",
		TitleName: "普卢里沃",
		TitleCode: "b_plourivo",
	}
}
