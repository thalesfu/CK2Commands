package dege

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 八邦BabangBarony struct {
	feud.BaseBarony
}

var BBabang八邦 feud.Barony = &八邦BabangBarony{}

func init() {
	f := BBabang八邦.(*八邦BabangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "babang",
		TitleName: "八邦",
		TitleCode: "b_babang",
	}
}
