package naxos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣托里尼SantoriniBarony struct {
	feud.BaseBarony
}

var BSantorini圣托里尼 feud.Barony = &圣托里尼SantoriniBarony{}

func init() {
    f := BSantorini圣托里尼.(*圣托里尼SantoriniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "santorini",
		TitleName: "圣托里尼",
		TitleCode: "b_santorini",
	}
}
