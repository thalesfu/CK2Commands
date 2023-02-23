package lingtsang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 麻呷MagaBarony struct {
	feud.BaseBarony
}

var BMaga麻呷 feud.Barony = &麻呷MagaBarony{}

func init() {
	f := BMaga麻呷.(*麻呷MagaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maga",
		TitleName: "麻呷",
		TitleCode: "b_maga",
	}
}
