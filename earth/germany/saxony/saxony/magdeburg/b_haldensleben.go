package magdeburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈尔登斯莱本HaldenslebenBarony struct {
	feud.BaseBarony
}

var BHaldensleben哈尔登斯莱本 feud.Barony = &哈尔登斯莱本HaldenslebenBarony{}

func init() {
	f := BHaldensleben哈尔登斯莱本.(*哈尔登斯莱本HaldenslebenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "haldensleben",
		TitleName: "哈尔登斯莱本",
		TitleCode: "b_haldensleben",
	}
}
