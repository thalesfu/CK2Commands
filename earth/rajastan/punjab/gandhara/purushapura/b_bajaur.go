package purushapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴焦尔BajaurBarony struct {
	feud.BaseBarony
}

var BBajaur巴焦尔 feud.Barony = &巴焦尔BajaurBarony{}

func init() {
	f := BBajaur巴焦尔.(*巴焦尔BajaurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bajaur",
		TitleName: "巴焦尔",
		TitleCode: "b_bajaur",
	}
}
