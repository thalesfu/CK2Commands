package jeddah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欧斯凡UsfanBarony struct {
	feud.BaseBarony
}

var BUsfan欧斯凡 feud.Barony = &欧斯凡UsfanBarony{}

func init() {
    f := BUsfan欧斯凡.(*欧斯凡UsfanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "usfan",
		TitleName: "欧斯凡",
		TitleCode: "b_usfan",
	}
}
