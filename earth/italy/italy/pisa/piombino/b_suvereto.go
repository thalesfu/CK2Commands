package piombino

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏韦雷托SuveretoBarony struct {
	feud.BaseBarony
}

var BSuvereto苏韦雷托 feud.Barony = &苏韦雷托SuveretoBarony{}

func init() {
    f := BSuvereto苏韦雷托.(*苏韦雷托SuveretoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "suvereto",
		TitleName: "苏韦雷托",
		TitleCode: "b_suvereto",
	}
}
