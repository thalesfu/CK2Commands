package hlynov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫雷诺夫HlynovBarony struct {
	feud.BaseBarony
}

var BHlynov赫雷诺夫 feud.Barony = &赫雷诺夫HlynovBarony{}

func init() {
    f := BHlynov赫雷诺夫.(*赫雷诺夫HlynovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hlynov",
		TitleName: "赫雷诺夫",
		TitleCode: "b_hlynov",
	}
}
