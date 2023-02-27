package mansia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼西斯克MansiyskBarony struct {
	feud.BaseBarony
}

var BMansiysk曼西斯克 feud.Barony = &曼西斯克MansiyskBarony{}

func init() {
    f := BMansiysk曼西斯克.(*曼西斯克MansiyskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mansiysk",
		TitleName: "曼西斯克",
		TitleCode: "b_mansiysk",
	}
}
