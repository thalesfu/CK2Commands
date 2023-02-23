package breisgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗赖堡FreiburgBarony struct {
	feud.BaseBarony
}

var BFreiburg弗赖堡 feud.Barony = &弗赖堡FreiburgBarony{}

func init() {
	f := BFreiburg弗赖堡.(*弗赖堡FreiburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "freiburg",
		TitleName: "弗赖堡",
		TitleCode: "b_freiburg",
	}
}
