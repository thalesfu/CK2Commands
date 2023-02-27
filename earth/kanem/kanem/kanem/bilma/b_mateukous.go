package bilma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马特库斯MateukousBarony struct {
	feud.BaseBarony
}

var BMateukous马特库斯 feud.Barony = &马特库斯MateukousBarony{}

func init() {
    f := BMateukous马特库斯.(*马特库斯MateukousBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mateukous",
		TitleName: "马特库斯",
		TitleCode: "b_mateukous",
	}
}
