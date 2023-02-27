package firenze

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佛罗伦萨FirenzeBarony struct {
	feud.BaseBarony
}

var BFirenze佛罗伦萨 feud.Barony = &佛罗伦萨FirenzeBarony{}

func init() {
    f := BFirenze佛罗伦萨.(*佛罗伦萨FirenzeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "firenze",
		TitleName: "佛罗伦萨",
		TitleCode: "b_firenze",
	}
}
