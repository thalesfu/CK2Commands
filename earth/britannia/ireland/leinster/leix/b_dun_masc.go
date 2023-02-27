package leix

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 敦马斯克Dun_mascBarony struct {
	feud.BaseBarony
}

var BDun_masc敦马斯克 feud.Barony = &敦马斯克Dun_mascBarony{}

func init() {
    f := BDun_masc敦马斯克.(*敦马斯克Dun_mascBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dun_masc",
		TitleName: "敦马斯克",
		TitleCode: "b_dun_masc",
	}
}
