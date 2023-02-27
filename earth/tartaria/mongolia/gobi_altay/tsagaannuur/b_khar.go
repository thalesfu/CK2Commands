package tsagaannuur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈尔KharBarony struct {
	feud.BaseBarony
}

var BKhar哈尔 feud.Barony = &哈尔KharBarony{}

func init() {
    f := BKhar哈尔.(*哈尔KharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khar",
		TitleName: "哈尔",
		TitleCode: "b_khar",
	}
}
