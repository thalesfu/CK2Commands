package trier

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科布伦茨CoblenzBarony struct {
	feud.BaseBarony
}

var BCoblenz科布伦茨 feud.Barony = &科布伦茨CoblenzBarony{}

func init() {
    f := BCoblenz科布伦茨.(*科布伦茨CoblenzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "coblenz",
		TitleName: "科布伦茨",
		TitleCode: "b_coblenz",
	}
}
