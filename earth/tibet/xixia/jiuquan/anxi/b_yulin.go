package anxi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 榆林YulinBarony struct {
	feud.BaseBarony
}

var BYulin榆林 feud.Barony = &榆林YulinBarony{}

func init() {
	f := BYulin榆林.(*榆林YulinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yulin",
		TitleName: "榆林",
		TitleCode: "b_yulin",
	}
}
