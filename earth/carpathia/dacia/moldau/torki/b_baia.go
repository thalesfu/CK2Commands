package torki

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴亚BaiaBarony struct {
	feud.BaseBarony
}

var BBaia巴亚 feud.Barony = &巴亚BaiaBarony{}

func init() {
	f := BBaia巴亚.(*巴亚BaiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baia",
		TitleName: "巴亚",
		TitleCode: "b_baia",
	}
}
