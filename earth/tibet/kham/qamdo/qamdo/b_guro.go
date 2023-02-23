package qamdo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 俄洛GuroBarony struct {
	feud.BaseBarony
}

var BGuro俄洛 feud.Barony = &俄洛GuroBarony{}

func init() {
	f := BGuro俄洛.(*俄洛GuroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guro",
		TitleName: "俄洛",
		TitleCode: "b_guro",
	}
}
