package lubusz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏伦钦SulecinBarony struct {
	feud.BaseBarony
}

var BSulecin苏伦钦 feud.Barony = &苏伦钦SulecinBarony{}

func init() {
    f := BSulecin苏伦钦.(*苏伦钦SulecinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sulecin",
		TitleName: "苏伦钦",
		TitleCode: "b_sulecin",
	}
}
