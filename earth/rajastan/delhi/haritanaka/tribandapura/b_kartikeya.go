package tribandapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦絺吉夜KartikeyaBarony struct {
	feud.BaseBarony
}

var BKartikeya迦絺吉夜 feud.Barony = &迦絺吉夜KartikeyaBarony{}

func init() {
    f := BKartikeya迦絺吉夜.(*迦絺吉夜KartikeyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kartikeya",
		TitleName: "迦絺吉夜",
		TitleCode: "b_kartikeya",
	}
}
