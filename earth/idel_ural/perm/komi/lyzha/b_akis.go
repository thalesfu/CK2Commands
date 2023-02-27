package lyzha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿基斯AkisBarony struct {
	feud.BaseBarony
}

var BAkis阿基斯 feud.Barony = &阿基斯AkisBarony{}

func init() {
    f := BAkis阿基斯.(*阿基斯AkisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akis",
		TitleName: "阿基斯",
		TitleCode: "b_akis",
	}
}
