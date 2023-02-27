package kara_khoja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 交河JiaoheBarony struct {
	feud.BaseBarony
}

var BJiaohe交河 feud.Barony = &交河JiaoheBarony{}

func init() {
    f := BJiaohe交河.(*交河JiaoheBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jiaohe",
		TitleName: "交河",
		TitleCode: "b_jiaohe",
	}
}
