package vastergotland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 林德霍尔门LindholmenBarony struct {
	feud.BaseBarony
}

var BLindholmen林德霍尔门 feud.Barony = &林德霍尔门LindholmenBarony{}

func init() {
    f := BLindholmen林德霍尔门.(*林德霍尔门LindholmenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lindholmen",
		TitleName: "林德霍尔门",
		TitleCode: "b_lindholmen",
	}
}
