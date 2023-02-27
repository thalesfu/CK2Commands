package urgench

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克孜尔阿吉尔Kizil_agirBarony struct {
	feud.BaseBarony
}

var BKizil_agir克孜尔阿吉尔 feud.Barony = &克孜尔阿吉尔Kizil_agirBarony{}

func init() {
    f := BKizil_agir克孜尔阿吉尔.(*克孜尔阿吉尔Kizil_agirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kizil_agir",
		TitleName: "克孜尔阿吉尔",
		TitleCode: "b_kizil_agir",
	}
}
