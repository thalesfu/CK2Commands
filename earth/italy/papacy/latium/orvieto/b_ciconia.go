package orvieto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奇科尼亚CiconiaBarony struct {
	feud.BaseBarony
}

var BCiconia奇科尼亚 feud.Barony = &奇科尼亚CiconiaBarony{}

func init() {
	f := BCiconia奇科尼亚.(*奇科尼亚CiconiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ciconia",
		TitleName: "奇科尼亚",
		TitleCode: "b_ciconia",
	}
}
