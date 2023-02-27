package konjikala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉奥克杰佩GokdepeBarony struct {
	feud.BaseBarony
}

var BGokdepe吉奥克杰佩 feud.Barony = &吉奥克杰佩GokdepeBarony{}

func init() {
    f := BGokdepe吉奥克杰佩.(*吉奥克杰佩GokdepeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gokdepe",
		TitleName: "吉奥克杰佩",
		TitleCode: "b_gokdepe",
	}
}
