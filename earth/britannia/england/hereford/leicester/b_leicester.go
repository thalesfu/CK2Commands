package leicester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱斯特LeicesterBarony struct {
	feud.BaseBarony
}

var BLeicester莱斯特 feud.Barony = &莱斯特LeicesterBarony{}

func init() {
	f := BLeicester莱斯特.(*莱斯特LeicesterBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "leicester",
		TitleName: "莱斯特",
		TitleCode: "b_leicester",
	}
}
