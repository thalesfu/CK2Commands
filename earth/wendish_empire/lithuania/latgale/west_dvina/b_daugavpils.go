package west_dvina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陶格夫匹尔斯DaugavpilsBarony struct {
	feud.BaseBarony
}

var BDaugavpils陶格夫匹尔斯 feud.Barony = &陶格夫匹尔斯DaugavpilsBarony{}

func init() {
    f := BDaugavpils陶格夫匹尔斯.(*陶格夫匹尔斯DaugavpilsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "daugavpils",
		TitleName: "陶格夫匹尔斯",
		TitleCode: "b_daugavpils",
	}
}
