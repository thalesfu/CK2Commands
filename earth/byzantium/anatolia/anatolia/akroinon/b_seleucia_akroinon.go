package akroinon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞琉西亚Seleucia_akroinonBarony struct {
	feud.BaseBarony
}

var BSeleucia_akroinon塞琉西亚 feud.Barony = &塞琉西亚Seleucia_akroinonBarony{}

func init() {
    f := BSeleucia_akroinon塞琉西亚.(*塞琉西亚Seleucia_akroinonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "seleucia_akroinon",
		TitleName: "塞琉西亚",
		TitleCode: "b_seleucia_akroinon",
	}
}
