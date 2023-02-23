package shirvan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙赫里亚尔ShakhriyarBarony struct {
	feud.BaseBarony
}

var BShakhriyar沙赫里亚尔 feud.Barony = &沙赫里亚尔ShakhriyarBarony{}

func init() {
	f := BShakhriyar沙赫里亚尔.(*沙赫里亚尔ShakhriyarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shakhriyar",
		TitleName: "沙赫里亚尔",
		TitleCode: "b_shakhriyar",
	}
}
