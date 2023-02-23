package dashhowuz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库巴达格GubadagBarony struct {
	feud.BaseBarony
}

var BGubadag库巴达格 feud.Barony = &库巴达格GubadagBarony{}

func init() {
	f := BGubadag库巴达格.(*库巴达格GubadagBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gubadag",
		TitleName: "库巴达格",
		TitleCode: "b_gubadag",
	}
}
