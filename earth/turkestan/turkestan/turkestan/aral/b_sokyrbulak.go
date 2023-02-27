package aral

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索克尔_布拉克SokyrbulakBarony struct {
	feud.BaseBarony
}

var BSokyrbulak索克尔_布拉克 feud.Barony = &索克尔_布拉克SokyrbulakBarony{}

func init() {
    f := BSokyrbulak索克尔_布拉克.(*索克尔_布拉克SokyrbulakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sokyrbulak",
		TitleName: "索克尔_布拉克",
		TitleCode: "b_sokyrbulak",
	}
}
