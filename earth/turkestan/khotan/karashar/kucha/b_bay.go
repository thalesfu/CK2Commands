package kucha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拜城BayBarony struct {
	feud.BaseBarony
}

var BBay拜城 feud.Barony = &拜城BayBarony{}

func init() {
	f := BBay拜城.(*拜城BayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bay",
		TitleName: "拜城",
		TitleCode: "b_bay",
	}
}
