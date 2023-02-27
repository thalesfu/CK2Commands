package bolzano

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃潘EppanBarony struct {
	feud.BaseBarony
}

var BEppan埃潘 feud.Barony = &埃潘EppanBarony{}

func init() {
    f := BEppan埃潘.(*埃潘EppanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eppan",
		TitleName: "埃潘",
		TitleCode: "b_eppan",
	}
}
