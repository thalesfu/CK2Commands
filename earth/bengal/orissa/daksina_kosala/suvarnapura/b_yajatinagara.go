package suvarnapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶阇底那揭罗YajatinagaraBarony struct {
	feud.BaseBarony
}

var BYajatinagara耶阇底那揭罗 feud.Barony = &耶阇底那揭罗YajatinagaraBarony{}

func init() {
    f := BYajatinagara耶阇底那揭罗.(*耶阇底那揭罗YajatinagaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yajatinagara",
		TitleName: "耶阇底那揭罗",
		TitleCode: "b_yajatinagara",
	}
}
