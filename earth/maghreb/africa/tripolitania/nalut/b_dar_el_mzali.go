package nalut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达尔姆扎利Dar_el_mzaliBarony struct {
	feud.BaseBarony
}

var BDar_el_mzali达尔姆扎利 feud.Barony = &达尔姆扎利Dar_el_mzaliBarony{}

func init() {
    f := BDar_el_mzali达尔姆扎利.(*达尔姆扎利Dar_el_mzaliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dar_el_mzali",
		TitleName: "达尔姆扎利",
		TitleCode: "b_dar_el_mzali",
	}
}
