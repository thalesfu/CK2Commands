package dorylaion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 日尔米亚GermiaBarony struct {
	feud.BaseBarony
}

var BGermia日尔米亚 feud.Barony = &日尔米亚GermiaBarony{}

func init() {
	f := BGermia日尔米亚.(*日尔米亚GermiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "germia",
		TitleName: "日尔米亚",
		TitleCode: "b_germia",
	}
}
