package qwivir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙赫尔古米斯SharequmisBarony struct {
	feud.BaseBarony
}

var BSharequmis沙赫尔古米斯 feud.Barony = &沙赫尔古米斯SharequmisBarony{}

func init() {
	f := BSharequmis沙赫尔古米斯.(*沙赫尔古米斯SharequmisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sharequmis",
		TitleName: "沙赫尔古米斯",
		TitleCode: "b_sharequmis",
	}
}
