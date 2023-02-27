package jaen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克萨达QuesadaBarony struct {
	feud.BaseBarony
}

var BQuesada克萨达 feud.Barony = &克萨达QuesadaBarony{}

func init() {
    f := BQuesada克萨达.(*克萨达QuesadaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "quesada",
		TitleName: "克萨达",
		TitleCode: "b_quesada",
	}
}
