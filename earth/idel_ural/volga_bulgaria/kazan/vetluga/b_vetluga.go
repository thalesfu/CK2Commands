package vetluga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦特卢加VetlugaBarony struct {
	feud.BaseBarony
}

var BVetluga韦特卢加 feud.Barony = &韦特卢加VetlugaBarony{}

func init() {
    f := BVetluga韦特卢加.(*韦特卢加VetlugaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vetluga",
		TitleName: "韦特卢加",
		TitleCode: "b_vetluga",
	}
}
