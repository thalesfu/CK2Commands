package don_portage

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛克LoqBarony struct {
	feud.BaseBarony
}

var BLoq洛克 feud.Barony = &洛克LoqBarony{}

func init() {
    f := BLoq洛克.(*洛克LoqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "loq",
		TitleName: "洛克",
		TitleCode: "b_loq",
	}
}
