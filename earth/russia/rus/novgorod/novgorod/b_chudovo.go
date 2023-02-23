package novgorod

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丘多沃ChudovoBarony struct {
	feud.BaseBarony
}

var BChudovo丘多沃 feud.Barony = &丘多沃ChudovoBarony{}

func init() {
	f := BChudovo丘多沃.(*丘多沃ChudovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chudovo",
		TitleName: "丘多沃",
		TitleCode: "b_chudovo",
	}
}
