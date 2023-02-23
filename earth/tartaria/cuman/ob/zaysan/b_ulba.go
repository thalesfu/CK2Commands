package zaysan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌尔巴UlbaBarony struct {
	feud.BaseBarony
}

var BUlba乌尔巴 feud.Barony = &乌尔巴UlbaBarony{}

func init() {
	f := BUlba乌尔巴.(*乌尔巴UlbaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ulba",
		TitleName: "乌尔巴",
		TitleCode: "b_ulba",
	}
}
