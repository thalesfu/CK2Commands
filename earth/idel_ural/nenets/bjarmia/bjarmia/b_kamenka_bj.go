package bjarmia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡缅卡Kamenka_bjBarony struct {
	feud.BaseBarony
}

var BKamenka_bj卡缅卡 feud.Barony = &卡缅卡Kamenka_bjBarony{}

func init() {
    f := BKamenka_bj卡缅卡.(*卡缅卡Kamenka_bjBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kamenka_bj",
		TitleName: "卡缅卡",
		TitleCode: "b_kamenka_bj",
	}
}
