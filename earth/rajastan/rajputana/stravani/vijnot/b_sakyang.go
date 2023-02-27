package vijnot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨姜SakyangBarony struct {
	feud.BaseBarony
}

var BSakyang萨姜 feud.Barony = &萨姜SakyangBarony{}

func init() {
    f := BSakyang萨姜.(*萨姜SakyangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sakyang",
		TitleName: "萨姜",
		TitleCode: "b_sakyang",
	}
}
