package ejin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 玉河YuheBarony struct {
	feud.BaseBarony
}

var BYuhe玉河 feud.Barony = &玉河YuheBarony{}

func init() {
	f := BYuhe玉河.(*玉河YuheBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yuhe",
		TitleName: "玉河",
		TitleCode: "b_yuhe",
	}
}
