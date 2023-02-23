package qarazhyrya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克孜勒套KyzyltauBarony struct {
	feud.BaseBarony
}

var BKyzyltau克孜勒套 feud.Barony = &克孜勒套KyzyltauBarony{}

func init() {
	f := BKyzyltau克孜勒套.(*克孜勒套KyzyltauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kyzyltau",
		TitleName: "克孜勒套",
		TitleCode: "b_kyzyltau",
	}
}
