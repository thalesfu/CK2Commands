package dorset

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙夫茨伯里ShaftesburyBarony struct {
	feud.BaseBarony
}

var BShaftesbury沙夫茨伯里 feud.Barony = &沙夫茨伯里ShaftesburyBarony{}

func init() {
    f := BShaftesbury沙夫茨伯里.(*沙夫茨伯里ShaftesburyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shaftesbury",
		TitleName: "沙夫茨伯里",
		TitleCode: "b_shaftesbury",
	}
}
