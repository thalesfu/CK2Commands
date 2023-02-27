package karin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡林KarinBarony struct {
	feud.BaseBarony
}

var BKarin卡林 feud.Barony = &卡林KarinBarony{}

func init() {
    f := BKarin卡林.(*卡林KarinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karin",
		TitleName: "卡林",
		TitleCode: "b_karin",
	}
}
