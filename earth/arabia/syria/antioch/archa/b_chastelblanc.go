package archa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布朗堡ChastelblancBarony struct {
	feud.BaseBarony
}

var BChastelblanc布朗堡 feud.Barony = &布朗堡ChastelblancBarony{}

func init() {
    f := BChastelblanc布朗堡.(*布朗堡ChastelblancBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chastelblanc",
		TitleName: "布朗堡",
		TitleCode: "b_chastelblanc",
	}
}
