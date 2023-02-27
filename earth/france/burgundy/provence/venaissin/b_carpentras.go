package venaissin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡庞特拉CarpentrasBarony struct {
	feud.BaseBarony
}

var BCarpentras卡庞特拉 feud.Barony = &卡庞特拉CarpentrasBarony{}

func init() {
    f := BCarpentras卡庞特拉.(*卡庞特拉CarpentrasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "carpentras",
		TitleName: "卡庞特拉",
		TitleCode: "b_carpentras",
	}
}
