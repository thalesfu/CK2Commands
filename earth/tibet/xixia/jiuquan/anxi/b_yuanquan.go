package anxi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 渊泉YuanquanBarony struct {
	feud.BaseBarony
}

var BYuanquan渊泉 feud.Barony = &渊泉YuanquanBarony{}

func init() {
	f := BYuanquan渊泉.(*渊泉YuanquanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yuanquan",
		TitleName: "渊泉",
		TitleCode: "b_yuanquan",
	}
}
