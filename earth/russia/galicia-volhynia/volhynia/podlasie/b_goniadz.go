package podlasie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈尼翁兹GoniadzBarony struct {
	feud.BaseBarony
}

var BGoniadz戈尼翁兹 feud.Barony = &戈尼翁兹GoniadzBarony{}

func init() {
    f := BGoniadz戈尼翁兹.(*戈尼翁兹GoniadzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "goniadz",
		TitleName: "戈尼翁兹",
		TitleCode: "b_goniadz",
	}
}
