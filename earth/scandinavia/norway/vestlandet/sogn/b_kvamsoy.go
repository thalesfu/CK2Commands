package sogn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克瓦姆岛KvamsoyBarony struct {
	feud.BaseBarony
}

var BKvamsoy克瓦姆岛 feud.Barony = &克瓦姆岛KvamsoyBarony{}

func init() {
	f := BKvamsoy克瓦姆岛.(*克瓦姆岛KvamsoyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kvamsoy",
		TitleName: "克瓦姆岛",
		TitleCode: "b_kvamsoy",
	}
}
