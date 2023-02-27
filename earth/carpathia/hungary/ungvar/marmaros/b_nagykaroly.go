package marmaros

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑙吉卡罗伊NagykarolyBarony struct {
	feud.BaseBarony
}

var BNagykaroly瑙吉卡罗伊 feud.Barony = &瑙吉卡罗伊NagykarolyBarony{}

func init() {
    f := BNagykaroly瑙吉卡罗伊.(*瑙吉卡罗伊NagykarolyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nagykaroly",
		TitleName: "瑙吉卡罗伊",
		TitleCode: "b_nagykaroly",
	}
}
