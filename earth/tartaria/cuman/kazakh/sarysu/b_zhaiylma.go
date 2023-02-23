package sarysu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎伊尔马ZhaiylmaBarony struct {
	feud.BaseBarony
}

var BZhaiylma扎伊尔马 feud.Barony = &扎伊尔马ZhaiylmaBarony{}

func init() {
	f := BZhaiylma扎伊尔马.(*扎伊尔马ZhaiylmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhaiylma",
		TitleName: "扎伊尔马",
		TitleCode: "b_zhaiylma",
	}
}
