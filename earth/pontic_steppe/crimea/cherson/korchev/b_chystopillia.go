package korchev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基斯托比利亚ChystopilliaBarony struct {
	feud.BaseBarony
}

var BChystopillia基斯托比利亚 feud.Barony = &基斯托比利亚ChystopilliaBarony{}

func init() {
    f := BChystopillia基斯托比利亚.(*基斯托比利亚ChystopilliaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chystopillia",
		TitleName: "基斯托比利亚",
		TitleCode: "b_chystopillia",
	}
}
