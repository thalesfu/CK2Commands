package torzhok

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡缅卡KamenkaBarony struct {
	feud.BaseBarony
}

var BKamenka卡缅卡 feud.Barony = &卡缅卡KamenkaBarony{}

func init() {
    f := BKamenka卡缅卡.(*卡缅卡KamenkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kamenka",
		TitleName: "卡缅卡",
		TitleCode: "b_kamenka",
	}
}
