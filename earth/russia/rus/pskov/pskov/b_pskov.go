package pskov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普斯科夫PskovBarony struct {
	feud.BaseBarony
}

var BPskov普斯科夫 feud.Barony = &普斯科夫PskovBarony{}

func init() {
    f := BPskov普斯科夫.(*普斯科夫PskovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pskov",
		TitleName: "普斯科夫",
		TitleCode: "b_pskov",
	}
}
