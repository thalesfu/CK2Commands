package tribandapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帝利畔荼城TribandapuraBarony struct {
	feud.BaseBarony
}

var BTribandapura帝利畔荼城 feud.Barony = &帝利畔荼城TribandapuraBarony{}

func init() {
    f := BTribandapura帝利畔荼城.(*帝利畔荼城TribandapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tribandapura",
		TitleName: "帝利畔荼城",
		TitleCode: "b_tribandapura",
	}
}
