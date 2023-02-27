package zyriane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏克孙SuksunBarony struct {
	feud.BaseBarony
}

var BSuksun苏克孙 feud.Barony = &苏克孙SuksunBarony{}

func init() {
    f := BSuksun苏克孙.(*苏克孙SuksunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "suksun",
		TitleName: "苏克孙",
		TitleCode: "b_suksun",
	}
}
