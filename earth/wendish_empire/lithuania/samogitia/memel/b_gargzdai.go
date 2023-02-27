package memel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加尔格兹代GargzdaiBarony struct {
	feud.BaseBarony
}

var BGargzdai加尔格兹代 feud.Barony = &加尔格兹代GargzdaiBarony{}

func init() {
    f := BGargzdai加尔格兹代.(*加尔格兹代GargzdaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gargzdai",
		TitleName: "加尔格兹代",
		TitleCode: "b_gargzdai",
	}
}
