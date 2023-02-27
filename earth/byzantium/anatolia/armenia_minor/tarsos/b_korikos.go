package tarsos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科里科斯KorikosBarony struct {
	feud.BaseBarony
}

var BKorikos科里科斯 feud.Barony = &科里科斯KorikosBarony{}

func init() {
    f := BKorikos科里科斯.(*科里科斯KorikosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "korikos",
		TitleName: "科里科斯",
		TitleCode: "b_korikos",
	}
}
