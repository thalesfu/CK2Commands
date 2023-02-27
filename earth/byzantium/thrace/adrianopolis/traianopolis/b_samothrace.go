package traianopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨莫色雷斯SamothraceBarony struct {
	feud.BaseBarony
}

var BSamothrace萨莫色雷斯 feud.Barony = &萨莫色雷斯SamothraceBarony{}

func init() {
    f := BSamothrace萨莫色雷斯.(*萨莫色雷斯SamothraceBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "samothrace",
		TitleName: "萨莫色雷斯",
		TitleCode: "b_samothrace",
	}
}
