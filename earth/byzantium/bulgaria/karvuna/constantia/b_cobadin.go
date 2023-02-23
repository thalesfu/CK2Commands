package constantia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科巴丁CobadinBarony struct {
	feud.BaseBarony
}

var BCobadin科巴丁 feud.Barony = &科巴丁CobadinBarony{}

func init() {
	f := BCobadin科巴丁.(*科巴丁CobadinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cobadin",
		TitleName: "科巴丁",
		TitleCode: "b_cobadin",
	}
}
