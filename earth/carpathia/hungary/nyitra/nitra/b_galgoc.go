package nitra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 高尔戈茨GalgocBarony struct {
	feud.BaseBarony
}

var BGalgoc高尔戈茨 feud.Barony = &高尔戈茨GalgocBarony{}

func init() {
	f := BGalgoc高尔戈茨.(*高尔戈茨GalgocBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "galgoc",
		TitleName: "高尔戈茨",
		TitleCode: "b_galgoc",
	}
}
