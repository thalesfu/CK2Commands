package medapata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勒沙婆提婆RishabhdeoBarony struct {
	feud.BaseBarony
}

var BRishabhdeo勒沙婆提婆 feud.Barony = &勒沙婆提婆RishabhdeoBarony{}

func init() {
	f := BRishabhdeo勒沙婆提婆.(*勒沙婆提婆RishabhdeoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rishabhdeo",
		TitleName: "勒沙婆提婆",
		TitleCode: "b_rishabhdeo",
	}
}
