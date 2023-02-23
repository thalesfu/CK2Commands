package senj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔洛巴格KarlobagBarony struct {
	feud.BaseBarony
}

var BKarlobag卡尔洛巴格 feud.Barony = &卡尔洛巴格KarlobagBarony{}

func init() {
	f := BKarlobag卡尔洛巴格.(*卡尔洛巴格KarlobagBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karlobag",
		TitleName: "卡尔洛巴格",
		TitleCode: "b_karlobag",
	}
}
