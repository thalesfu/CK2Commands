package suvarnapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦罗汉持KalahandiBarony struct {
	feud.BaseBarony
}

var BKalahandi迦罗汉持 feud.Barony = &迦罗汉持KalahandiBarony{}

func init() {
    f := BKalahandi迦罗汉持.(*迦罗汉持KalahandiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalahandi",
		TitleName: "迦罗汉持",
		TitleCode: "b_kalahandi",
	}
}
