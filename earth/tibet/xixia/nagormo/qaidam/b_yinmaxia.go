package qaidam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 饮马峡YinmaxiaBarony struct {
	feud.BaseBarony
}

var BYinmaxia饮马峡 feud.Barony = &饮马峡YinmaxiaBarony{}

func init() {
	f := BYinmaxia饮马峡.(*饮马峡YinmaxiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yinmaxia",
		TitleName: "饮马峡",
		TitleCode: "b_yinmaxia",
	}
}
