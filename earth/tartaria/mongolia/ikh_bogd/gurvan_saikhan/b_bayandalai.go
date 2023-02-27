package gurvan_saikhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴彦达赖BayandalaiBarony struct {
	feud.BaseBarony
}

var BBayandalai巴彦达赖 feud.Barony = &巴彦达赖BayandalaiBarony{}

func init() {
    f := BBayandalai巴彦达赖.(*巴彦达赖BayandalaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bayandalai",
		TitleName: "巴彦达赖",
		TitleCode: "b_bayandalai",
	}
}
