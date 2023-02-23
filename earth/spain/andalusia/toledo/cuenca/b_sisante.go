package cuenca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西桑特SisanteBarony struct {
	feud.BaseBarony
}

var BSisante西桑特 feud.Barony = &西桑特SisanteBarony{}

func init() {
	f := BSisante西桑特.(*西桑特SisanteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sisante",
		TitleName: "西桑特",
		TitleCode: "b_sisante",
	}
}
