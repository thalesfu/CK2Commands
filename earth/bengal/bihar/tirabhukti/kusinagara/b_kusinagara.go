package kusinagara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拘尸那揭罗KusinagaraBarony struct {
	feud.BaseBarony
}

var BKusinagara拘尸那揭罗 feud.Barony = &拘尸那揭罗KusinagaraBarony{}

func init() {
    f := BKusinagara拘尸那揭罗.(*拘尸那揭罗KusinagaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kusinagara",
		TitleName: "拘尸那揭罗",
		TitleCode: "b_kusinagara",
	}
}
