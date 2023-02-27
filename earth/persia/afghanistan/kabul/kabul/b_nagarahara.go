package kabul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那揭罗曷NagaraharaBarony struct {
	feud.BaseBarony
}

var BNagarahara那揭罗曷 feud.Barony = &那揭罗曷NagaraharaBarony{}

func init() {
    f := BNagarahara那揭罗曷.(*那揭罗曷NagaraharaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nagarahara",
		TitleName: "那揭罗曷",
		TitleCode: "b_nagarahara",
	}
}
