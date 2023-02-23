package burgos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡里翁CarrionBarony struct {
	feud.BaseBarony
}

var BCarrion卡里翁 feud.Barony = &卡里翁CarrionBarony{}

func init() {
	f := BCarrion卡里翁.(*卡里翁CarrionBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "carrion",
		TitleName: "卡里翁",
		TitleCode: "b_carrion",
	}
}
