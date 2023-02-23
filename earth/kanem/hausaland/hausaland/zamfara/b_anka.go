package zamfara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安卡AnkaBarony struct {
	feud.BaseBarony
}

var BAnka安卡 feud.Barony = &安卡AnkaBarony{}

func init() {
	f := BAnka安卡.(*安卡AnkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anka",
		TitleName: "安卡",
		TitleCode: "b_anka",
	}
}
