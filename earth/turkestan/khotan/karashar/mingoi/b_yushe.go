package mingoi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 玉舍YusheBarony struct {
	feud.BaseBarony
}

var BYushe玉舍 feud.Barony = &玉舍YusheBarony{}

func init() {
    f := BYushe玉舍.(*玉舍YusheBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yushe",
		TitleName: "玉舍",
		TitleCode: "b_yushe",
	}
}
