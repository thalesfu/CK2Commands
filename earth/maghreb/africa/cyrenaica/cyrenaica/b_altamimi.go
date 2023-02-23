package cyrenaica

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰米米AltamimiBarony struct {
	feud.BaseBarony
}

var BAltamimi泰米米 feud.Barony = &泰米米AltamimiBarony{}

func init() {
	f := BAltamimi泰米米.(*泰米米AltamimiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "altamimi",
		TitleName: "泰米米",
		TitleCode: "b_altamimi",
	}
}
