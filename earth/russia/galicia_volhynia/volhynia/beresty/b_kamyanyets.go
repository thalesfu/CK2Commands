package beresty

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡梅涅茨KamyanyetsBarony struct {
	feud.BaseBarony
}

var BKamyanyets卡梅涅茨 feud.Barony = &卡梅涅茨KamyanyetsBarony{}

func init() {
    f := BKamyanyets卡梅涅茨.(*卡梅涅茨KamyanyetsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kamyanyets",
		TitleName: "卡梅涅茨",
		TitleCode: "b_kamyanyets",
	}
}
