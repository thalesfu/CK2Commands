package lepiel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙尔科夫希纳SharkawshchynaBarony struct {
	feud.BaseBarony
}

var BSharkawshchyna沙尔科夫希纳 feud.Barony = &沙尔科夫希纳SharkawshchynaBarony{}

func init() {
    f := BSharkawshchyna沙尔科夫希纳.(*沙尔科夫希纳SharkawshchynaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sharkawshchyna",
		TitleName: "沙尔科夫希纳",
		TitleCode: "b_sharkawshchyna",
	}
}
