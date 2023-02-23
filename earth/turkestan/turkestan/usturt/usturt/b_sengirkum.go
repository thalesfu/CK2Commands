package usturt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 先吉尔_库姆SengirkumBarony struct {
	feud.BaseBarony
}

var BSengirkum先吉尔_库姆 feud.Barony = &先吉尔_库姆SengirkumBarony{}

func init() {
	f := BSengirkum先吉尔_库姆.(*先吉尔_库姆SengirkumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sengirkum",
		TitleName: "先吉尔_库姆",
		TitleCode: "b_sengirkum",
	}
}
