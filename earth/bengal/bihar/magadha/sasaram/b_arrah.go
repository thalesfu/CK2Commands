package sasaram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿罗ArrahBarony struct {
	feud.BaseBarony
}

var BArrah阿罗 feud.Barony = &阿罗ArrahBarony{}

func init() {
    f := BArrah阿罗.(*阿罗ArrahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arrah",
		TitleName: "阿罗",
		TitleCode: "b_arrah",
	}
}
