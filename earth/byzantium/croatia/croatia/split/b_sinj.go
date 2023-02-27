package split

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西尼SinjBarony struct {
	feud.BaseBarony
}

var BSinj西尼 feud.Barony = &西尼SinjBarony{}

func init() {
    f := BSinj西尼.(*西尼SinjBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sinj",
		TitleName: "西尼",
		TitleCode: "b_sinj",
	}
}
