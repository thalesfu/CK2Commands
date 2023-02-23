package bannu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 汉古HanguBarony struct {
	feud.BaseBarony
}

var BHangu汉古 feud.Barony = &汉古HanguBarony{}

func init() {
	f := BHangu汉古.(*汉古HanguBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hangu",
		TitleName: "汉古",
		TitleCode: "b_hangu",
	}
}
