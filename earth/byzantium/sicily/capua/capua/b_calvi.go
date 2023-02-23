package capua

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔维CalviBarony struct {
	feud.BaseBarony
}

var BCalvi卡尔维 feud.Barony = &卡尔维CalviBarony{}

func init() {
	f := BCalvi卡尔维.(*卡尔维CalviBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "calvi",
		TitleName: "卡尔维",
		TitleCode: "b_calvi",
	}
}
