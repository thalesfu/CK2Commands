package zaysan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斋桑ZaysanBarony struct {
	feud.BaseBarony
}

var BZaysan斋桑 feud.Barony = &斋桑ZaysanBarony{}

func init() {
    f := BZaysan斋桑.(*斋桑ZaysanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zaysan",
		TitleName: "斋桑",
		TitleCode: "b_zaysan",
	}
}
