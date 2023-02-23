package perigord

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博纳吉耶BonaguilBarony struct {
	feud.BaseBarony
}

var BBonaguil博纳吉耶 feud.Barony = &博纳吉耶BonaguilBarony{}

func init() {
	f := BBonaguil博纳吉耶.(*博纳吉耶BonaguilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bonaguil",
		TitleName: "博纳吉耶",
		TitleCode: "b_bonaguil",
	}
}
