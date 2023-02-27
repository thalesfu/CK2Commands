package mallorca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马纳科尔ManacorBarony struct {
	feud.BaseBarony
}

var BManacor马纳科尔 feud.Barony = &马纳科尔ManacorBarony{}

func init() {
    f := BManacor马纳科尔.(*马纳科尔ManacorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "manacor",
		TitleName: "马纳科尔",
		TitleCode: "b_manacor",
	}
}
