package beirut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尤里涅JournieBarony struct {
	feud.BaseBarony
}

var BJournie尤里涅 feud.Barony = &尤里涅JournieBarony{}

func init() {
	f := BJournie尤里涅.(*尤里涅JournieBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "journie",
		TitleName: "尤里涅",
		TitleCode: "b_journie",
	}
}
