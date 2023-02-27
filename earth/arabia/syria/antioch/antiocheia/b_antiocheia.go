package antiocheia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安条克AntiocheiaBarony struct {
	feud.BaseBarony
}

var BAntiocheia安条克 feud.Barony = &安条克AntiocheiaBarony{}

func init() {
    f := BAntiocheia安条克.(*安条克AntiocheiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "antiocheia",
		TitleName: "安条克",
		TitleCode: "b_antiocheia",
	}
}
