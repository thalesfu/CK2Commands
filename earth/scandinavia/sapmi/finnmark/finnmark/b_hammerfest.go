package finnmark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈默弗斯特HammerfestBarony struct {
	feud.BaseBarony
}

var BHammerfest哈默弗斯特 feud.Barony = &哈默弗斯特HammerfestBarony{}

func init() {
    f := BHammerfest哈默弗斯特.(*哈默弗斯特HammerfestBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hammerfest",
		TitleName: "哈默弗斯特",
		TitleCode: "b_hammerfest",
	}
}
