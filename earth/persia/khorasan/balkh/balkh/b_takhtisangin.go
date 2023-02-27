package balkh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔赫季桑金TakhtisanginBarony struct {
	feud.BaseBarony
}

var BTakhtisangin塔赫季桑金 feud.Barony = &塔赫季桑金TakhtisanginBarony{}

func init() {
    f := BTakhtisangin塔赫季桑金.(*塔赫季桑金TakhtisanginBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "takhtisangin",
		TitleName: "塔赫季桑金",
		TitleCode: "b_takhtisangin",
	}
}
