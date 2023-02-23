package ancona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷卡纳蒂RecanatiBarony struct {
	feud.BaseBarony
}

var BRecanati雷卡纳蒂 feud.Barony = &雷卡纳蒂RecanatiBarony{}

func init() {
	f := BRecanati雷卡纳蒂.(*雷卡纳蒂RecanatiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "recanati",
		TitleName: "雷卡纳蒂",
		TitleCode: "b_recanati",
	}
}
