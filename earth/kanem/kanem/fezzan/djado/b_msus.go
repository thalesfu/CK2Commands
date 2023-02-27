package djado

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 姆苏斯MsusBarony struct {
	feud.BaseBarony
}

var BMsus姆苏斯 feud.Barony = &姆苏斯MsusBarony{}

func init() {
    f := BMsus姆苏斯.(*姆苏斯MsusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "msus",
		TitleName: "姆苏斯",
		TitleCode: "b_msus",
	}
}
