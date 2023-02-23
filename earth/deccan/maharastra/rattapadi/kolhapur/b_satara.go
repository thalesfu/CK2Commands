package kolhapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨达拉SataraBarony struct {
	feud.BaseBarony
}

var BSatara萨达拉 feud.Barony = &萨达拉SataraBarony{}

func init() {
	f := BSatara萨达拉.(*萨达拉SataraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "satara",
		TitleName: "萨达拉",
		TitleCode: "b_satara",
	}
}
