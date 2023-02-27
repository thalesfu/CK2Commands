package lisboa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴塔利亚BatalhaBarony struct {
	feud.BaseBarony
}

var BBatalha巴塔利亚 feud.Barony = &巴塔利亚BatalhaBarony{}

func init() {
    f := BBatalha巴塔利亚.(*巴塔利亚BatalhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "batalha",
		TitleName: "巴塔利亚",
		TitleCode: "b_batalha",
	}
}
