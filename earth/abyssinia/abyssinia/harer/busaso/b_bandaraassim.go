package busaso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 班达卡西姆BandaraassimBarony struct {
	feud.BaseBarony
}

var BBandaraassim班达卡西姆 feud.Barony = &班达卡西姆BandaraassimBarony{}

func init() {
	f := BBandaraassim班达卡西姆.(*班达卡西姆BandaraassimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bandaraassim",
		TitleName: "班达卡西姆",
		TitleCode: "b_bandaraassim",
	}
}
