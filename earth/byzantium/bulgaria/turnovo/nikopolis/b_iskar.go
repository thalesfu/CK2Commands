package nikopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊斯克尔IskarBarony struct {
	feud.BaseBarony
}

var BIskar伊斯克尔 feud.Barony = &伊斯克尔IskarBarony{}

func init() {
    f := BIskar伊斯克尔.(*伊斯克尔IskarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "iskar",
		TitleName: "伊斯克尔",
		TitleCode: "b_iskar",
	}
}
