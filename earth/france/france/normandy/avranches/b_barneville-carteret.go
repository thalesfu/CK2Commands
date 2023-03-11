package avranches

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴尔纳维尔卡尔特雷Barneville_carteretBarony struct {
	feud.BaseBarony
}

var BBarneville_carteret巴尔纳维尔卡尔特雷 feud.Barony = &巴尔纳维尔卡尔特雷Barneville_carteretBarony{}

func init() {
    f := BBarneville_carteret巴尔纳维尔卡尔特雷.(*巴尔纳维尔卡尔特雷Barneville_carteretBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barneville_carteret",
		TitleName: "巴尔纳维尔卡尔特雷",
		TitleCode: "b_barneville-carteret",
	}
}
