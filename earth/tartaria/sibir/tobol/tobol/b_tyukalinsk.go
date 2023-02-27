package tobol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 秋卡林斯克TyukalinskBarony struct {
	feud.BaseBarony
}

var BTyukalinsk秋卡林斯克 feud.Barony = &秋卡林斯克TyukalinskBarony{}

func init() {
    f := BTyukalinsk秋卡林斯克.(*秋卡林斯克TyukalinskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tyukalinsk",
		TitleName: "秋卡林斯克",
		TitleCode: "b_tyukalinsk",
	}
}
