package khantia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉涅什DjineshBarony struct {
	feud.BaseBarony
}

var BDjinesh吉涅什 feud.Barony = &吉涅什DjineshBarony{}

func init() {
    f := BDjinesh吉涅什.(*吉涅什DjineshBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "djinesh",
		TitleName: "吉涅什",
		TitleCode: "b_djinesh",
	}
}
